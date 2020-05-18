package endpoints

import (
	"github.com/CherryDock/CherryDock/api/docker/actions"
	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"net/http"
)

func handleSingleAction(w http.ResponseWriter, r *http.Request, singleAction actions.Action) {
	id := r.FormValue("id")
	if id != "" {
		success := actions.ActionSingleContainer(singleAction, id)
		if success == true {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Fail to execute action, id not exists", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Fail to execute action, id parameter is missing", http.StatusBadRequest)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.StartContainer)
}

func stop(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.StopContainer)
}

func restart(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.RestartContainer)
}

func remove(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.RemoveContainer)
}

func kill(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.KillContainer)
}

func pause(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.PauseContainer)
}

func unpause(w http.ResponseWriter, r *http.Request) {
	handleSingleAction(w, r, actions.UnpauseContainer)
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

func monitor(w http.ResponseWriter, r *http.Request) {
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
