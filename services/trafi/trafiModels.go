package trafi

type RealtimeData struct {
	VehicleId        string   `json:"vehicleId"`
	DepartsInMinutes int      `json:"departsInMinutes"`
	Destination      string   `json:"destination"`
	Tags             []string `json:"tags"`
}

type Schedule struct {
	TrackName string         `json:"trackName"`
	TrackId   string         `json:"TrackId"`
	StopName  string         `json:"stopName"`
	StopId    string         `json:"stopId"`
	Realtime  []RealtimeData `json:"realtime"`
}
