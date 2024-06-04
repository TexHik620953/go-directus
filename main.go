package main

import (
	"go-directus/directus"
	"log"
)

func main() {
	directusApi, err := directus.New("http://localhost:8055", "hH7xIXYzonjs_HUpdhNFLgKuwYIbWkPe")
	if err != nil {
		log.Fatal(err)
	}

	server, err := directusApi.GameServersCollectionAccessor.ReadAll().
		Where(`name == "!23"`).
		Include("*, moderators.*, project.*").
		First()
	if err != nil {
		log.Fatal(err)
	}

	err = directusApi.SaveChanges()
	if err != nil {
		log.Fatal(err)
	}

}
