package routes

import (
	"encoding/json"
	"fmt"
	"hatvservice/config"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)

type Channel struct {
	Id          string `json:"id"`
	ContentType string `json:"contentType"`
	StreamUrl   string `json:"streamUrl"`
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	config := config.GetConfig()
	vars := mux.Vars(r)
	channelId := vars["channelId"]
	var channel Channel

	for _, chn := range config.Channels {
		if chn.Id == channelId {
			streamUrl, err := getStreamUrl(&chn)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			channel = Channel{
				Id:          chn.Id,
				StreamUrl:   *streamUrl,
				ContentType: chn.ContentType,
			}
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if (channel == Channel{}) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(channel)
	}
}

func getStreamUrl(channelConfig *config.ConfigChannel) (*string, error) {
	resp, err := http.Get(channelConfig.Url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	strBody := string(body)
	streamUrl := gjson.Get(strBody, channelConfig.StreamUrlPath).String()

	return &streamUrl, nil
}
