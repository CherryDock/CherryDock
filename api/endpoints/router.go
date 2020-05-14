package endpoints

import (
	"github.com/CherryDock/CherryDock/api/jwt"
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Get jwt token route
	router.HandleFunc("/token", jwt.GetToken)

	// Monitoring routes
	router.Handle("/api/monitor/containers-info", jwt.CheckToken(http.HandlerFunc(getContainersInfoJson)))
	router.Handle("/api/monitor/logs", jwt.CheckToken(http.HandlerFunc(getLogs)))
	router.Handle("/api/monitor/stats", jwt.CheckToken(http.HandlerFunc(monitorAll)))
	router.Handle("/api/monitor/stat", jwt.CheckToken(http.HandlerFunc(monitorSingle)))

	// Actions routes
	router.Handle("/api/action/stop-all", jwt.CheckToken(http.HandlerFunc(stopAll)))
	router.Handle("/api/action/start-all", jwt.CheckToken(http.HandlerFunc(startAll)))
	router.Handle("/api/action/start", jwt.CheckToken(http.HandlerFunc(startSingle)))
	router.Handle("/api/action/stop", jwt.CheckToken(http.HandlerFunc(stopSingle)))
	router.Handle("/api/action/restart", jwt.CheckToken(http.HandlerFunc(restartSingle)))
	router.Handle("/api/action/remove", jwt.CheckToken(http.HandlerFunc(removeSingle)))

	return router
}
