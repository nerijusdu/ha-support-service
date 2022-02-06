package trafi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRealtimeSchedule(scheduleId string, stopId string, trackId string) (*Schedule, error) {
	resp, err := http.Get("https://web.trafi.com/api/times/vilnius/realtime?scheduleId=" + scheduleId + "&trackId=" + trackId + "&stopId=" + stopId)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	schedule := &Schedule{}
	err = json.Unmarshal(body, schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
