package directus

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/google/uuid"
)

type DirectusResponse[T any] struct {
	Data   T `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type trackingRef struct {
	Original        IDirectusObject
	Actual          IDirectusObject
	OwnerCollection IDirectusCollectionAccessor
}

func (h trackingRef) delta() map[string]any {
	return h.Actual.(IDirectusObject).Diff(h.Original.(IDirectusObject))
}

type DirectusApi struct {
	directusUrl *url.URL
	token       string

	errLogger  *log.Logger
	infoLogger *log.Logger

	DirectusActivityCollectionAccessor      *DirectusCollectionAccessor[int, DirectusActivity]
	DirectusDashboardsCollectionAccessor    *DirectusCollectionAccessor[uuid.UUID, DirectusDashboards]
	DirectusExtensionsCollectionAccessor    *DirectusCollectionAccessor[uuid.UUID, DirectusExtensions]
	DirectusFieldsCollectionAccessor        *DirectusCollectionAccessor[int, DirectusFields]
	DirectusFilesCollectionAccessor         *DirectusCollectionAccessor[uuid.UUID, DirectusFiles]
	DirectusFlowsCollectionAccessor         *DirectusCollectionAccessor[uuid.UUID, DirectusFlows]
	DirectusFoldersCollectionAccessor       *DirectusCollectionAccessor[uuid.UUID, DirectusFolders]
	DirectusNotificationsCollectionAccessor *DirectusCollectionAccessor[int, DirectusNotifications]
	DirectusOperationsCollectionAccessor    *DirectusCollectionAccessor[uuid.UUID, DirectusOperations]
	DirectusPanelsCollectionAccessor        *DirectusCollectionAccessor[uuid.UUID, DirectusPanels]
	DirectusPermissionsCollectionAccessor   *DirectusCollectionAccessor[int, DirectusPermissions]
	DirectusPresetsCollectionAccessor       *DirectusCollectionAccessor[int, DirectusPresets]
	DirectusRelationsCollectionAccessor     *DirectusCollectionAccessor[int, DirectusRelations]
	DirectusRevisionsCollectionAccessor     *DirectusCollectionAccessor[int, DirectusRevisions]
	DirectusRolesCollectionAccessor         *DirectusCollectionAccessor[uuid.UUID, DirectusRoles]
	DirectusSettingsCollectionAccessor      *DirectusCollectionAccessor[int, DirectusSettings]
	DirectusSharesCollectionAccessor        *DirectusCollectionAccessor[uuid.UUID, DirectusShares]
	DirectusTranslationsCollectionAccessor  *DirectusCollectionAccessor[uuid.UUID, DirectusTranslations]
	DirectusUsersCollectionAccessor         *DirectusCollectionAccessor[uuid.UUID, DirectusUsers]
	DirectusVersionsCollectionAccessor      *DirectusCollectionAccessor[uuid.UUID, DirectusVersions]
	DirectusWebhooksCollectionAccessor      *DirectusCollectionAccessor[int, DirectusWebhooks]
	LocationCollectionAccessor              *DirectusCollectionAccessor[uuid.UUID, Location]
	ProductCollectionAccessor               *DirectusCollectionAccessor[uuid.UUID, Product]
	PromocodeCollectionAccessor             *DirectusCollectionAccessor[uuid.UUID, Promocode]
	ProxyServerCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, ProxyServer]
	SlotCollectionAccessor                  *DirectusCollectionAccessor[uuid.UUID, Slot]
	TransactionCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, Transaction]

	collectionsAccessors map[string]IDirectusCollectionAccessor
}

func New(addr, token string) (*DirectusApi, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	h := &DirectusApi{
		directusUrl: u,
		token:       token,
		errLogger:   log.New(os.Stdout, "[DIRECTUS-API][ERROR]\t", log.Ltime),
		infoLogger:  log.New(os.Stdout, "[DIRECTUS-API][INFO]\t", log.Ltime),
	}
	err = h.PingDirectus()
	if err != nil {
		return nil, err
	}

	h.DirectusActivityCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusActivity](h, "directus_activity")
	h.DirectusDashboardsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusDashboards](h, "directus_dashboards")
	h.DirectusExtensionsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusExtensions](h, "directus_extensions")
	h.DirectusFieldsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusFields](h, "directus_fields")
	h.DirectusFilesCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusFiles](h, "directus_files")
	h.DirectusFlowsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusFlows](h, "directus_flows")
	h.DirectusFoldersCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusFolders](h, "directus_folders")
	h.DirectusNotificationsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusNotifications](h, "directus_notifications")
	h.DirectusOperationsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusOperations](h, "directus_operations")
	h.DirectusPanelsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusPanels](h, "directus_panels")
	h.DirectusPermissionsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusPermissions](h, "directus_permissions")
	h.DirectusPresetsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusPresets](h, "directus_presets")
	h.DirectusRelationsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusRelations](h, "directus_relations")
	h.DirectusRevisionsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusRevisions](h, "directus_revisions")
	h.DirectusRolesCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusRoles](h, "directus_roles")
	h.DirectusSettingsCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusSettings](h, "directus_settings")
	h.DirectusSharesCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusShares](h, "directus_shares")
	h.DirectusTranslationsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusTranslations](h, "directus_translations")
	h.DirectusUsersCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusUsers](h, "directus_users")
	h.DirectusVersionsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusVersions](h, "directus_versions")
	h.DirectusWebhooksCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusWebhooks](h, "directus_webhooks")
	h.LocationCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, Location](h, "location")
	h.ProductCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, Product](h, "product")
	h.PromocodeCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, Promocode](h, "promocode")
	h.ProxyServerCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, ProxyServer](h, "proxy_server")
	h.SlotCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, Slot](h, "slot")
	h.TransactionCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, Transaction](h, "transaction")

	h.collectionsAccessors = map[string]IDirectusCollectionAccessor{
		"directus_activity":      h.DirectusActivityCollectionAccessor,
		"directus_dashboards":    h.DirectusDashboardsCollectionAccessor,
		"directus_extensions":    h.DirectusExtensionsCollectionAccessor,
		"directus_fields":        h.DirectusFieldsCollectionAccessor,
		"directus_files":         h.DirectusFilesCollectionAccessor,
		"directus_flows":         h.DirectusFlowsCollectionAccessor,
		"directus_folders":       h.DirectusFoldersCollectionAccessor,
		"directus_notifications": h.DirectusNotificationsCollectionAccessor,
		"directus_operations":    h.DirectusOperationsCollectionAccessor,
		"directus_panels":        h.DirectusPanelsCollectionAccessor,
		"directus_permissions":   h.DirectusPermissionsCollectionAccessor,
		"directus_presets":       h.DirectusPresetsCollectionAccessor,
		"directus_relations":     h.DirectusRelationsCollectionAccessor,
		"directus_revisions":     h.DirectusRevisionsCollectionAccessor,
		"directus_roles":         h.DirectusRolesCollectionAccessor,
		"directus_settings":      h.DirectusSettingsCollectionAccessor,
		"directus_shares":        h.DirectusSharesCollectionAccessor,
		"directus_translations":  h.DirectusTranslationsCollectionAccessor,
		"directus_users":         h.DirectusUsersCollectionAccessor,
		"directus_versions":      h.DirectusVersionsCollectionAccessor,
		"directus_webhooks":      h.DirectusWebhooksCollectionAccessor,
		"location":               h.LocationCollectionAccessor,
		"product":                h.ProductCollectionAccessor,
		"promocode":              h.PromocodeCollectionAccessor,
		"proxy_server":           h.ProxyServerCollectionAccessor,
		"slot":                   h.SlotCollectionAccessor,
		"transaction":            h.TransactionCollectionAccessor,
	}

	return h, nil
}

func (h *DirectusApi) PingDirectus() error {
	addr := *h.directusUrl
	addr.Path = path.Join(addr.Path, "/server/ping")

	resp, err := http.Get(addr.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		h.errLogger.Printf("Directus ping failed: %s\n", err.Error())
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func NewDirectusCollectionAccessor[K string | uuid.UUID | int, V IDirectusObject](api *DirectusApi, collectionName string) *DirectusCollectionAccessor[K, V] {
	api.infoLogger.Printf("Created collection accessor for %s\n", collectionName)
	return &DirectusCollectionAccessor[K, V]{
		api:            api,
		collectionName: collectionName,
	}
}

func key2String[K string | uuid.UUID | int](key K) string {
	switch any(key).(type) {
	case string:
		return any(key).(string)
	case uuid.UUID:
		return any(key).(uuid.UUID).String()
	case uint32:
		return strconv.FormatUint(uint64(any(key).(uint32)), 10)
	}
	// This code should not be reachable
	panic("How did you get there?")
}
