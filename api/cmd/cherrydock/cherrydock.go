package main

import (
	"github.com/CherryDock/CherryDock/api/configuration"
	"github.com/CherryDock/CherryDock/api/database"
	"github.com/CherryDock/CherryDock/api/http"
	"github.com/CherryDock/CherryDock/api/scheduler"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	config := configuration.LoadConfig(pwd + "/config.yml")

	// Init bolt embedded db
	database.DbClient = &database.Client{}
	database.DbClient.Init()

	// Start monitoring thread
	scheduler.ScheduleMonitoring()

	http.StartServer(config)
}
