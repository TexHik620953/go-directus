package main

import (
	"fmt"
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
		Include("*, *.*").
		First()
	if err != nil {
		log.Fatal(err)
	}
	*server.Project.Annotation = "Srakotan"

	err = directusApi.SaveChanges()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(server)
}
