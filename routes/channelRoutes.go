package routes

import (
	"encoding/json"
	"fmt"
	"haservice/config"
	"haservice/services/channels"
	"net/http"

	"github.com/gorilla/mux"
)

func getChannel(w http.ResponseWriter, r *http.Request) {
	config := config.GetConfig()
	vars := mux.Vars(r)
	channelId := vars["channelId"]
	var channel channels.Channel

	for _, chn := range config.Channels {
		if chn.Id == channelId {
			streamUrl, err := channels.GetStreamUrl(&chn)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			channel = channels.Channel{
				Id:          chn.Id,
				StreamUrl:   *streamUrl,
				ContentType: chn.ContentType,
			}
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if (channel == channels.Channel{}) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(channel)
	}
}
