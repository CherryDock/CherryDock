package main

import (
	"github.com/CherryDock/CherryDock/configuration"
	"github.com/CherryDock/CherryDock/endpoints"
	"log"
	"net/http"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	config := configuration.LoadConfig(pwd + "/config.yml")
	server := config.Server
	router := endpoints.Routing()

	log.Printf("Start http server at %v:%v", server.Host, server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, router))
}
