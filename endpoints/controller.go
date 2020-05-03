package endpoints

import (
	"github.com/Fszta/DockerMonitoring/docker/actions"
	"github.com/Fszta/DockerMonitoring/docker/monitoring"
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/monitor/memory-stats", getMemoryJson)
	router.HandleFunc("/api/monitor/containers-info", getContainersInfoJson)
	router.HandleFunc("/api/monitor/logs", getLogs)
	router.HandleFunc("/api/action/stop-all", stopAll)
	router.HandleFunc("/api/action/start-all", startAll)
	router.HandleFunc("/api/action/start", StartSingle)
	router.HandleFunc("/api/action/stop", StopSingle)
	router.HandleFunc("/api/action/restart", RestartSingle)

	return router
}

func StartSingle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		success := actions.ActionSingleContainer(actions.StartContainer, id)
		if success == true {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Fail to start container, id not exists", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Fail to stop container, id parameter is missing", http.StatusBadRequest)
	}
}

func StopSingle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		success := actions.ActionSingleContainer(actions.StopContainer, id)

		if success == true {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Fail to stop container, id not exists", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Fail to stop container, id parameter is missing", http.StatusBadRequest)
	}
}

func RestartSingle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		success := actions.ActionSingleContainer(actions.RestartContainer, id)

		if success == true {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Fail to restart container, id not exists", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Fail to restart container, id parameter is missing", http.StatusBadRequest)
	}
}

func startAll(w http.ResponseWriter, r *http.Request) {
	states := actions.ActionOnAllContainer(actions.StartContainer, true)
	w.Header().Set("content-type", "application/json")
	w.Write(states)
}

func stopAll(w http.ResponseWriter, r *http.Request) {
	states := actions.ActionOnAllContainer(actions.StopContainer, false)
	w.Header().Set("content-type", "application/json")
	w.Write(states)
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

func getLogs(w http.ResponseWriter, r *http.Request) {
	containerdId := r.FormValue("id")
	logs := monitoring.RetrieveLogs(containerdId)

	w.Write(logs)
}
