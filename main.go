package main

import (
	"fmt"
	"go-directus/directus"
	"log"

	"github.com/google/uuid"
)

func main() {
	directusApi, err := directus.New("http://localhost:8055")
	if err != nil {
		log.Fatal(err)
	}
	gameServers := directus.NewDirectusCollection[uuid.UUID, directus.GameServers](directusApi, "GameServers", "hH7xIXYzonjs_HUpdhNFLgKuwYIbWkPe")

	//id := "00996200-b4d6-4fcf-b6ac-ad6ce25139cc"
	server, err := gameServers.ReadAll().
		Where(`name == "123"`).
		Include("*, *.*").
		First()
	if err != nil {
		log.Fatal(err)
	}
	*server.Address = "123"

	err = gameServers.SaveChanges()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(server)
}
