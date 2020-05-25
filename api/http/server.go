package http

import (
	"github.com/CherryDock/CherryDock/api/configuration"
	"github.com/CherryDock/CherryDock/api/endpoints"
	"log"
	"net/http"
)

func StartServer(config *configuration.Configuration) {
	server := config.Server
	router := endpoints.Routing()
	log.Printf("Start http server at %v:%v", server.Host, server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, router))
}
