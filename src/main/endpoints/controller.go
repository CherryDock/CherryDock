package endpoints

import (
	"github.com/Fszta/DockerMonitoring/src/main/docker/actions"
	"github.com/Fszta/DockerMonitoring/src/main/docker/monitoring"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/memory-stats", getMemoryJson)
	router.HandleFunc("/api/containers-info",getContainersInfoJson)
	router.HandleFunc("/api/start-container/{id}", launchContainer)
	router.HandleFunc("/api/start-container/{id}", launchContainer)
	router.HandleFunc("/api/stop-container/{id}", stopContainer)

	return router
}

func test(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	log.Print(vars["val"])
}

func getMemoryJson(w http.ResponseWriter, r *http.Request) {
	memoryStats := monitoring.GetMemoryStats()
	w.Header().Set("content-type", "application/json")
	w.Write(memoryStats)
}

func getContainersInfoJson(w http.ResponseWriter, r *http.Request) {
	containersInfo := monitoring.GetContainersInfo(true)
	w.Header().Set("content-type","application/json")
	w.Write(containersInfo)
}

func launchContainer(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	actions.StartContainer(parameters["id"])
	log.Printf("Start container %s",parameters["id"])
}

func stopContainer(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	actions.StopContainer(parameters["id"])
	log.Printf("Start container %s",parameters["id"])
}
