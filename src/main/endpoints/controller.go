package endpoints

import (
	"github.com/Fszta/DockerMonitoring/src/main/docker/actions"
	"github.com/Fszta/DockerMonitoring/src/main/docker/monitoring"
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/memory-stats", getMemoryJson)
	router.HandleFunc("/api/containers-info", getContainersInfoJson)
	router.HandleFunc("/api/action/{action}/{id}", handleAction)

	return router
}

func getMemoryJson(w http.ResponseWriter, r *http.Request) {
	memoryStats := monitoring.GetMemoryStats()
	w.Header().Set("content-type", "application/json")
	w.Write(memoryStats)
}

func getContainersInfoJson(w http.ResponseWriter, r *http.Request) {
	containersInfo := monitoring.GetContainersInfo(true)
	w.Header().Set("content-type", "application/json")
	w.Write(containersInfo)
}

func handleAction(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	action := parameters["action"]
	id := parameters["id"]

	// Handle single container action
	actions.Handle(action, id)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
