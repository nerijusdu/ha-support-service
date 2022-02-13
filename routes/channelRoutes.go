package routes

import (
	"fmt"
	"haservice/config"
	"haservice/services/channels"
	"haservice/utils"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func getChannelConfig(r *http.Request) *config.ConfigChannel {
	config := config.GetConfig()
	vars := mux.Vars(r)
	channelId := vars["channelId"]

	for _, chn := range config.Channels {
		if chn.Id == channelId {
			return &chn
		}
	}

	return nil
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	channelConfig := getChannelConfig(r)

	if channelConfig == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	streamUrl, err := channels.GetStreamUrl(channelConfig)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	channel := channels.Channel{
		Id:          channelConfig.Id,
		StreamUrl:   *streamUrl,
		ContentType: channelConfig.ContentType,
	}

	utils.WriteJson(w, channel)
}

func getChannelProgram(w http.ResponseWriter, r *http.Request) {
	channelConfig := getChannelConfig(r)

	if channelConfig == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	program, err := channels.GetProgram(channelConfig)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/channels/program.html"))
	tmpl.Execute(w, ProgramTemplateData{
		Content:    *program,
		Stylesheet: channelConfig.ProgramStylesheet,
		ScrollTo:   channelConfig.ProgramScrollTo,
	})
}

type ProgramTemplateData struct {
	Content    string
	Stylesheet string
	ScrollTo   string
}
