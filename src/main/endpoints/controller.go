package endpoints

import (
	"github.com/gorilla/mux"
	"net/http"
	"test/src/main/docker/monitoring"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/memory-stats", getMemoryJson)
	router.HandleFunc("/api/containers-info",getContainersInfoJson)

	return router
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