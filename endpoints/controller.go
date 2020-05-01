package endpoints

import (
	"github.com/Fszta/DockerMonitoring/docker/actions"
	"github.com/Fszta/DockerMonitoring/docker/monitoring"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/memory-stats", getMemoryJson)
	router.HandleFunc("/api/containers-info", getContainersInfoJson)
	router.HandleFunc("/api/action/", handleAction)

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
	/*parameters := mux.Vars(r)
	action := parameters["action"]
	id := parameters["id"]*/
	action := r.FormValue("action")
	id := r.FormValue("id")
	// Handle single container action
	log.Printf(action, id)
	actions.Handle(action, id)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
