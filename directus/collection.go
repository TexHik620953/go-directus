package directus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/google/uuid"
)

type IDirectusCollectionAccessor interface {
	patch(object map[string]any, id string) error
}

type DirectusCollectionAccessor[K string | uuid.UUID | int, V IDirectusObject] struct {
	IDirectusCollectionAccessor
	api            *DirectusApi
	collectionName string
}

func (h *DirectusCollectionAccessor[K, V]) LoadById(id K) (*V, error) {
	addr := *h.api.directusUrl
	addr.Path = path.Join(addr.Path, fmt.Sprintf("/items/%s/%s", h.collectionName, key2String(id)))
	req, err := http.NewRequest("GET", addr.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", h.api.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	item := DirectusResponse[*V]{}
	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	if item.Errors != nil {
		msg := ""
		if len(item.Errors) != 0 {
			msg = item.Errors[0].Message
		}
		return nil, fmt.Errorf(msg)
	}
	h.api.add2Track(item.Data)
	return item.Data, nil
}

// FILTERING STREAM

// Readonly
type CollectionQuery[K string | uuid.UUID | int, V IDirectusObject] struct {
	Collection    *DirectusCollectionAccessor[K, V]
	customHeaders map[string]string

	whereFilters   []string
	fieldSelectors []string
}

func (h *DirectusCollectionAccessor[K, V]) ReadAll() *CollectionQuery[K, V] {
	return &CollectionQuery[K, V]{
		Collection:     h,
		whereFilters:   []string{},
		customHeaders:  map[string]string{},
		fieldSelectors: []string{},
	}
}

func (h *CollectionQuery[K, V]) Where(comparator ...string) *CollectionQuery[K, V] {
	h.whereFilters = append(h.whereFilters, comparator...)
	return h
}
func (h *CollectionQuery[K, V]) Include(selector ...string) *CollectionQuery[K, V] {
	for _, s := range selector {
		s = strings.ReplaceAll(s, " ", "")
		f := strings.Split(s, ",")
		h.fieldSelectors = append(h.fieldSelectors, f...)
	}
	return h
}

// Service
func (h *CollectionQuery[K, V]) WithCustomHeader(key, value string) *CollectionQuery[K, V] {
	h.customHeaders[key] = value
	return h
}
func (h *CollectionQuery[K, V]) ToSlice() ([]*V, error) {
	addr := *h.Collection.api.directusUrl
	addr.Path = path.Join(addr.Path, fmt.Sprintf("/items/%s", h.Collection.collectionName))
	q := addr.Query()

	filter, err := h.buildWhereFilters()
	if err != nil {
		return nil, err
	}

	q.Add("filter", filter)
	q.Add("fields", h.buildSelectors())
	addr.RawQuery = q.Encode()
	url := addr.String()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Collection.api.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	item := DirectusResponse[[]*V]{}
	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	if item.Errors != nil {
		msg := ""
		if len(item.Errors) != 0 {
			msg = item.Errors[0].Message
		}
		return nil, fmt.Errorf(msg)
	}

	for _, e := range item.Data {
		h.Collection.api.add2Track(e)
	}
	return item.Data, nil
}
func (h *CollectionQuery[K, V]) First() (*V, error) {
	addr := *h.Collection.api.directusUrl
	addr.Path = path.Join(addr.Path, fmt.Sprintf("/items/%s", h.Collection.collectionName))
	q := addr.Query()

	filter, err := h.buildWhereFilters()
	if err != nil {
		return nil, err
	}

	q.Add("filter", filter)
	q.Add("fields", h.buildSelectors())
	q.Add("limit", "1")
	addr.RawQuery = q.Encode()
	url := addr.String()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Collection.api.token))
	for k, v := range h.customHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	item := DirectusResponse[[]*V]{}

	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	if item.Errors != nil {
		msg := ""
		if len(item.Errors) != 0 {
			msg = item.Errors[0].Message
		}
		return nil, fmt.Errorf(msg)
	}
	if len(item.Data) == 0 {
		return nil, fmt.Errorf("Directus returned empty collection")
	}

	obj := item.Data[0]
	h.Collection.api.add2Track(obj)
	return obj, nil
}

func (h *DirectusCollectionAccessor[K, V]) patch(object map[string]any, id string) error {
	addr := *h.api.directusUrl
	addr.Path = path.Join(addr.Path, fmt.Sprintf("/items/%s/%s", h.collectionName, id))

	body, err := json.Marshal(object)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", addr.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.api.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	item := DirectusResponse[*V]{}
	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		return err
	}
	if item.Errors != nil {
		msg := ""
		if len(item.Errors) != 0 {
			msg = item.Errors[0].Message
		}
		return fmt.Errorf(msg)
	}
	return nil
}
