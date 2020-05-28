package http

import (
	"log"
	"net/http"

	"github.com/CherryDock/CherryDock/api/configuration"
	"github.com/CherryDock/CherryDock/api/endpoints"
	"github.com/rs/cors"
)

func StartServer(config *configuration.Configuration) {
	server := config.Server
	router := endpoints.Routing()
	handler := cors.Default().Handler(router)
	log.Printf("Start http server at %v:%v", server.Host, server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, handler))
}
