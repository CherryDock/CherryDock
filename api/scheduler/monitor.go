package scheduler

import (
	"github.com/CherryDock/CherryDock/api/database"
	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"log"
	"time"
)

func ScheduleMonitoring() {
	db := database.DbClient
	ticker := time.NewTicker(10 * time.Second)
	log.Println("start monitoring containers")
	go func() {
		for t := range ticker.C {
			_ = t
			monitoringInfo := monitoring.GlobalMonitoring()
			db.AddMonitoringInfo(monitoringInfo)
		}
	}()
}
