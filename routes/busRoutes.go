package routes

import (
	"encoding/json"
	"fmt"
	"haservice/services/trafi"
	"net/http"
)

type NextBusResponse struct {
	Data []trafi.RealtimeData `json:"data"`
}

func getNextBus(w http.ResponseWriter, r *http.Request) {
	scheduleId := r.URL.Query().Get("scheduleId")
	stopId := r.URL.Query().Get("stopId")
	trackId := r.URL.Query().Get("trackId")

	if scheduleId == "" || stopId == "" || trackId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	schedule, err := trafi.GetRealtimeSchedule(scheduleId, stopId, trackId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&NextBusResponse{Data: schedule.Realtime})
}
