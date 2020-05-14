package endpoints

import (
	"github.com/CherryDock/CherryDock/api/docker/actions"
	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"net/http"
)

func startSingle(w http.ResponseWriter, r *http.Request) {
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

func stopSingle(w http.ResponseWriter, r *http.Request) {
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

func restartSingle(w http.ResponseWriter, r *http.Request) {
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

func removeSingle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		success := actions.ActionSingleContainer(actions.RemoveContainer, id)

		if success == true {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Fail to remove container, id not exists", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Fail to remove container, id parameter is missing", http.StatusBadRequest)
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

func monitorSingle(w http.ResponseWriter, r *http.Request) {
	containerdId := r.FormValue("id")
	stats := monitoring.SingleMonitoring(containerdId)
	w.Header().Set("content-type", "application/json")
	w.Write(stats)
}

func monitorAll(w http.ResponseWriter, r *http.Request) {
	globalStats := monitoring.GlobalMonitoring()
	w.Header().Set("content-type", "application/json")
	w.Write(globalStats)
}

func getContainersInfoJson(w http.ResponseWriter, r *http.Request) {
	containersInfo := monitoring.GetContainersInfo(true)
	w.Header().Set("content-type", "application/json")
	w.Write(containersInfo)
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	containerdId := r.FormValue("id")
	w.Header().Set("content-type", "application/json")

	if containerdId == "" {
		http.Error(w, "Fail to retrieve logs, id parameter is missing", http.StatusBadRequest)
	} else {
		logs, succeed := monitoring.RetrieveLogs(containerdId)
		if succeed == false {
			http.Error(w, "Fail to retrieve container logs, id not exists", http.StatusNotFound)
		} else {
			w.Write(logs)
		}
	}
}
