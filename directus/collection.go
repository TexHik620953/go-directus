package directus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type trackingRef struct {
	Original IDirectusObject
	Actual   IDirectusObject
}

func (h trackingRef) delta() map[string]any {
	return h.Actual.(IDirectusObject).Diff(h.Original.(IDirectusObject))
}

type DirectusCollection[K string | uuid.UUID | int, V IDirectusObject] struct {
	api            *DirectusApi
	collectionName string
	token          string

	trackingObjects map[IDirectusObject]trackingRef
}

func (h *DirectusCollection[K, V]) key2String(key K) string {
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
func (h *DirectusCollection[K, V]) add2Track(val *V) bool {
	objects := (*val).Track()
	for _, obj := range objects {
		_, exists := h.trackingObjects[obj]
		if !exists {
			obj_copy := (obj).DeepCopy()
			ref := trackingRef{
				Original: obj_copy,
				Actual:   obj,
			}
			h.trackingObjects[obj] = ref
		}
	}
	return false
}

func (h *DirectusCollection[K, V]) LoadById(id K) (*V, error) {
	addr := *h.api.directusUrl
	addr.Path = path.Join(addr.Path, fmt.Sprintf("/items/%s/%s", h.collectionName, h.key2String(id)))
	req, err := http.NewRequest("GET", addr.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", h.token)

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
	h.add2Track(item.Data)
	return item.Data, nil
}

func (h *DirectusCollection[K, V]) AddTrackingReference(val IDirectusObject) {
	//TODO
}

func (h *DirectusCollection[K, V]) SaveChanges() error {
	for _, obj := range h.trackingObjects {
		diff := obj.delta()
		if diff != nil {
			u, err := h.patch(diff, obj.Original.GetId())
			if err != nil {
				return err
			}
			fmt.Println(u)
		}
	}
	return nil
}

// FILTERING STREAM

// Readonly
type CollectionQuery[K string | uuid.UUID | int, V IDirectusObject] struct {
	Collection    *DirectusCollection[K, V]
	customHeaders map[string]string

	whereFilters   []string
	fieldSelectors []string
}

func (h *DirectusCollection[K, V]) ReadAll() *CollectionQuery[K, V] {
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Collection.token))

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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Collection.token))
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
	h.Collection.add2Track(obj)
	return obj, nil
}

func (h *DirectusCollection[K, V]) patch(object map[string]any, id string) (*V, error) {
	addr := *h.api.directusUrl
	addr.Path = path.Join(addr.Path, fmt.Sprintf("/items/%s/%s", h.collectionName, id))

	body, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", addr.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.token))

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
	return item.Data, nil
}
