package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/health", getHealth)
	router.HandleFunc("/api/channel/{channelId}", getChannel)
}

type HealthStatus struct {
	Healthy bool `json:"healthy"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&HealthStatus{Healthy: true})
}
