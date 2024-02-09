package directus

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/google/uuid"
)

var (
	UnexpectedStatusCode = errors.New("Unexpected status code")
)

type DirectusResponse[T any] struct {
	Data   T `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type DirectusApi struct {
	directusUrl *url.URL
}

func New(addr string) (*DirectusApi, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	h := &DirectusApi{
		directusUrl: u,
	}
	err = h.PingDirectus()
	if err != nil {
		return nil, err
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
		return fmt.Errorf("%w: unexpected status code: %d", UnexpectedStatusCode, resp.StatusCode)
	}
	return nil
}

func NewDirectusCollection[K string | uuid.UUID | int, V IDirectusObject](api *DirectusApi, collectionName, token string) *DirectusCollection[K, V] {
	return &DirectusCollection[K, V]{
		api:             api,
		collectionName:  collectionName,
		token:           token,
		trackingObjects: map[IDirectusObject]trackingRef{},
	}
}
