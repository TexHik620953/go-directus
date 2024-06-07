package directus

import (
	"log"
	"sync"
	"time"
)

type DirectusAccessContext struct {
	trackingObjects      map[IDirectusObject]trackingRef
	trackingObjectsMutex sync.Mutex
	api                  *DirectusApi
}

func (h *DirectusApi) NewDirectusAccessContext() *DirectusAccessContext {
	return &DirectusAccessContext{
		trackingObjects: map[IDirectusObject]trackingRef{},
		api:             h,
	}
}

func (h *DirectusAccessContext) add2Track(val any) bool {
	h.trackingObjectsMutex.Lock()
	defer h.trackingObjectsMutex.Unlock()
	objects := val.(IDirectusObject).Track()
	objects = append(objects, val.(IDirectusObject))
	for _, obj := range objects {
		_, exists := h.trackingObjects[obj]
		if !exists {
			ownerCollection, exists := h.api.collectionsAccessors[obj.CollectionName()]
			if !exists {
				log.Fatalf("Collection accessor for object: %s not exists in map", obj.CollectionName())
			}
			h.api.infoLogger.Printf("Added tracking reference for object of type [%s]\n", obj.CollectionName())
			obj_copy := obj.DeepCopy()
			ref := trackingRef{
				Original:        obj_copy,
				Actual:          obj,
				OwnerCollection: ownerCollection,
			}
			h.trackingObjects[obj] = ref
		}
	}
	return false
}

func (h *DirectusAccessContext) SaveChanges() error {
	h.trackingObjectsMutex.Lock()
	defer h.trackingObjectsMutex.Unlock()
	affectedObjects := 0
	startTime := time.Now()
	for _, obj := range h.trackingObjects {
		diff := obj.delta()
		if diff != nil {
			cas := obj.OwnerCollection
			err := cas.patch(diff, obj.Original.GetId())
			if err != nil {
				h.api.errLogger.Printf("Failed to save changes for object of type [%s]: %s\n", obj.Original.CollectionName(), err.Error())
				return err
			}
			affectedObjects++
		}
	}
	for io := range h.trackingObjects {
		delete(h.trackingObjects, io)
	}
	deltaTime := time.Since(startTime)
	h.api.infoLogger.Printf("Changes saved, affected [%d] objects, %s\n", affectedObjects, deltaTime)
	return nil
}

func (h *DirectusAccessContext) Clear() {
	h.trackingObjectsMutex.Lock()
	defer h.trackingObjectsMutex.Unlock()
	for io := range h.trackingObjects {
		delete(h.trackingObjects, io)
	}
}
