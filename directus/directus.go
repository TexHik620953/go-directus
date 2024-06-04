package directus

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

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
	if h.Original == nil {
		return h.Actual.(IDirectusObject).Map()
	}

	return h.Actual.(IDirectusObject).Diff(h.Original.(IDirectusObject))
}

type DirectusApi struct {
	directusUrl *url.URL
	token       string

	errLogger  *log.Logger
	infoLogger *log.Logger

	trackingObjects map[IDirectusObject]trackingRef

	CheckEventsCollectionAccessor             *DirectusCollectionAccessor[uuid.UUID, CheckEvents]
	CheckMessagesCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, CheckMessages]
	DirectusActivityCollectionAccessor        *DirectusCollectionAccessor[int, DirectusActivity]
	DirectusDashboardsCollectionAccessor      *DirectusCollectionAccessor[uuid.UUID, DirectusDashboards]
	DirectusFieldsCollectionAccessor          *DirectusCollectionAccessor[int, DirectusFields]
	DirectusFilesCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, DirectusFiles]
	DirectusFlowsCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, DirectusFlows]
	DirectusFoldersCollectionAccessor         *DirectusCollectionAccessor[uuid.UUID, DirectusFolders]
	DirectusNotificationsCollectionAccessor   *DirectusCollectionAccessor[int, DirectusNotifications]
	DirectusOperationsCollectionAccessor      *DirectusCollectionAccessor[uuid.UUID, DirectusOperations]
	DirectusPanelsCollectionAccessor          *DirectusCollectionAccessor[uuid.UUID, DirectusPanels]
	DirectusPermissionsCollectionAccessor     *DirectusCollectionAccessor[int, DirectusPermissions]
	DirectusPresetsCollectionAccessor         *DirectusCollectionAccessor[int, DirectusPresets]
	DirectusRelationsCollectionAccessor       *DirectusCollectionAccessor[int, DirectusRelations]
	DirectusRevisionsCollectionAccessor       *DirectusCollectionAccessor[int, DirectusRevisions]
	DirectusRolesCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, DirectusRoles]
	DirectusSettingsCollectionAccessor        *DirectusCollectionAccessor[int, DirectusSettings]
	DirectusSharesCollectionAccessor          *DirectusCollectionAccessor[uuid.UUID, DirectusShares]
	DirectusTranslationsCollectionAccessor    *DirectusCollectionAccessor[uuid.UUID, DirectusTranslations]
	DirectusUsersCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, DirectusUsers]
	DirectusVersionsCollectionAccessor        *DirectusCollectionAccessor[uuid.UUID, DirectusVersions]
	DirectusWebhooksCollectionAccessor        *DirectusCollectionAccessor[int, DirectusWebhooks]
	ExternalFilesCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, ExternalFiles]
	GameProjectsCollectionAccessor            *DirectusCollectionAccessor[uuid.UUID, GameProjects]
	GameServersCollectionAccessor             *DirectusCollectionAccessor[uuid.UUID, GameServers]
	GameServersModeratorsCollectionAccessor   *DirectusCollectionAccessor[int, GameServersModerators]
	KnownIpsCollectionAccessor                *DirectusCollectionAccessor[uuid.UUID, KnownIps]
	PlayerBansCollectionAccessor              *DirectusCollectionAccessor[uuid.UUID, PlayerBans]
	PlayerBansExternalFilesCollectionAccessor *DirectusCollectionAccessor[int, PlayerBansExternalFiles]
	PlayerChecksCollectionAccessor            *DirectusCollectionAccessor[uuid.UUID, PlayerChecks]
	PlayerEventsCollectionAccessor            *DirectusCollectionAccessor[uuid.UUID, PlayerEvents]
	PlayerReportsCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, PlayerReports]
	ServerChatMessagesCollectionAccessor      *DirectusCollectionAccessor[uuid.UUID, ServerChatMessages]
	ServerPlayersCollectionAccessor           *DirectusCollectionAccessor[uuid.UUID, ServerPlayers]
	SteamUsersCollectionAccessor              *DirectusCollectionAccessor[uuid.UUID, SteamUsers]
	TestCollectionAccessor                    *DirectusCollectionAccessor[uuid.UUID, Test]

	collectionsAccessors map[string]IDirectusCollectionAccessor
}

func New(addr, token string) (*DirectusApi, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	h := &DirectusApi{
		directusUrl:     u,
		token:           token,
		trackingObjects: map[IDirectusObject]trackingRef{},
		errLogger:       log.New(os.Stdout, "[DIRECTUS-API][ERROR]\t", log.Ltime),
		infoLogger:      log.New(os.Stdout, "[DIRECTUS-API][INFO]\t", log.Ltime),
	}
	err = h.PingDirectus()
	if err != nil {
		return nil, err
	}

	h.CheckEventsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, CheckEvents](h, "CheckEvents")
	h.CheckMessagesCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, CheckMessages](h, "CheckMessages")
	h.DirectusActivityCollectionAccessor = NewDirectusCollectionAccessor[int, DirectusActivity](h, "directus_activity")
	h.DirectusDashboardsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, DirectusDashboards](h, "directus_dashboards")
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
	h.ExternalFilesCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, ExternalFiles](h, "ExternalFiles")
	h.GameProjectsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, GameProjects](h, "GameProjects")
	h.GameServersCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, GameServers](h, "GameServers")
	h.GameServersModeratorsCollectionAccessor = NewDirectusCollectionAccessor[int, GameServersModerators](h, "GameServers_Moderators")
	h.KnownIpsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, KnownIps](h, "KnownIps")
	h.PlayerBansCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, PlayerBans](h, "PlayerBans")
	h.PlayerBansExternalFilesCollectionAccessor = NewDirectusCollectionAccessor[int, PlayerBansExternalFiles](h, "PlayerBans_ExternalFiles")
	h.PlayerChecksCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, PlayerChecks](h, "PlayerChecks")
	h.PlayerEventsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, PlayerEvents](h, "PlayerEvents")
	h.PlayerReportsCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, PlayerReports](h, "PlayerReports")
	h.ServerChatMessagesCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, ServerChatMessages](h, "ServerChatMessages")
	h.ServerPlayersCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, ServerPlayers](h, "ServerPlayers")
	h.SteamUsersCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, SteamUsers](h, "SteamUsers")
	h.TestCollectionAccessor = NewDirectusCollectionAccessor[uuid.UUID, Test](h, "test")

	h.collectionsAccessors = map[string]IDirectusCollectionAccessor{
		"CheckEvents":              h.CheckEventsCollectionAccessor,
		"CheckMessages":            h.CheckMessagesCollectionAccessor,
		"directus_activity":        h.DirectusActivityCollectionAccessor,
		"directus_dashboards":      h.DirectusDashboardsCollectionAccessor,
		"directus_fields":          h.DirectusFieldsCollectionAccessor,
		"directus_files":           h.DirectusFilesCollectionAccessor,
		"directus_flows":           h.DirectusFlowsCollectionAccessor,
		"directus_folders":         h.DirectusFoldersCollectionAccessor,
		"directus_notifications":   h.DirectusNotificationsCollectionAccessor,
		"directus_operations":      h.DirectusOperationsCollectionAccessor,
		"directus_panels":          h.DirectusPanelsCollectionAccessor,
		"directus_permissions":     h.DirectusPermissionsCollectionAccessor,
		"directus_presets":         h.DirectusPresetsCollectionAccessor,
		"directus_relations":       h.DirectusRelationsCollectionAccessor,
		"directus_revisions":       h.DirectusRevisionsCollectionAccessor,
		"directus_roles":           h.DirectusRolesCollectionAccessor,
		"directus_settings":        h.DirectusSettingsCollectionAccessor,
		"directus_shares":          h.DirectusSharesCollectionAccessor,
		"directus_translations":    h.DirectusTranslationsCollectionAccessor,
		"directus_users":           h.DirectusUsersCollectionAccessor,
		"directus_versions":        h.DirectusVersionsCollectionAccessor,
		"directus_webhooks":        h.DirectusWebhooksCollectionAccessor,
		"ExternalFiles":            h.ExternalFilesCollectionAccessor,
		"GameProjects":             h.GameProjectsCollectionAccessor,
		"GameServers":              h.GameServersCollectionAccessor,
		"GameServers_Moderators":   h.GameServersModeratorsCollectionAccessor,
		"KnownIps":                 h.KnownIpsCollectionAccessor,
		"PlayerBans":               h.PlayerBansCollectionAccessor,
		"PlayerBans_ExternalFiles": h.PlayerBansExternalFilesCollectionAccessor,
		"PlayerChecks":             h.PlayerChecksCollectionAccessor,
		"PlayerEvents":             h.PlayerEventsCollectionAccessor,
		"PlayerReports":            h.PlayerReportsCollectionAccessor,
		"ServerChatMessages":       h.ServerChatMessagesCollectionAccessor,
		"ServerPlayers":            h.ServerPlayersCollectionAccessor,
		"SteamUsers":               h.SteamUsersCollectionAccessor,
		"test":                     h.TestCollectionAccessor,
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

func (h *DirectusApi) add2Track(val any) {
	objects := val.(IDirectusObject).Track()
	objects = append(objects, val.(IDirectusObject))
	for _, obj := range objects {
		_, exists := h.trackingObjects[obj]
		if !exists {
			ownerCollection, exists := h.collectionsAccessors[obj.CollectionName()]
			if !exists {
				log.Fatalf("Collection accessor for object: %s not exists in map", obj.CollectionName())
			}
			h.infoLogger.Printf("Added tracking reference for object of type [%s]\n", obj.CollectionName())
			obj_copy := obj.DeepCopy()
			ref := trackingRef{
				Original:        obj_copy,
				Actual:          obj,
				OwnerCollection: ownerCollection,
			}
			h.trackingObjects[obj] = ref
		}
	}
}

func (h *DirectusApi) AddTrackingReference(val any) {
	objects := val.(IDirectusObject).Track()
	objects = append(objects, val.(IDirectusObject))
	for _, obj := range objects {
		_, exists := h.trackingObjects[obj]
		if !exists {
			ownerCollection, exists := h.collectionsAccessors[obj.CollectionName()]
			if !exists {
				log.Fatalf("Collection accessor for object: %s not exists in map", obj.CollectionName())
			}
			h.infoLogger.Printf("Added tracking reference for object of type [%s]\n", obj.CollectionName())
			ref := trackingRef{
				Actual:          obj,
				OwnerCollection: ownerCollection,
			}
			h.trackingObjects[obj] = ref
		}
	}
}

func (h *DirectusApi) SaveChanges() error {
	affectedObjects := 0
	startTime := time.Now()
	for _, obj := range h.trackingObjects {
		diff := obj.delta()
		if diff != nil {
			cas := obj.OwnerCollection
			err := cas.patch(diff, obj.Actual.GetId())
			if err != nil {
				h.errLogger.Printf("Failed to save changes for object of type [%s]: %s\n", obj.Original.CollectionName(), err.Error())
				return err
			}
			affectedObjects++
		}
	}
	deltaTime := time.Since(startTime)
	h.infoLogger.Printf("Changes saved, affected [%d] objects, %s\n", affectedObjects, deltaTime)
	return nil
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
