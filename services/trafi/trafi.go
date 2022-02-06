package trafi

import (
	"haservice/utils"
)

func GetRealtimeSchedule(scheduleId string, stopId string, trackId string) (*RealtimeSchedule, error) {
	schedule := &RealtimeSchedule{}
	err := utils.GetJson(
		"https://web.trafi.com/api/times/vilnius/realtime?scheduleId="+scheduleId+"&trackId="+trackId+"&stopId="+stopId,
		schedule,
	)

	return schedule, err
}
