package channels

import "html/template"

type Channel struct {
	Id          string `json:"id"`
	ContentType string `json:"contentType"`
	StreamUrl   string `json:"streamUrl"`
}

type ProgramTemplateData struct {
	Host       string
	Id         string
	Content    template.HTML
	Data       []ProgramItem
	Stylesheet string
	ScrollTo   string
}

type LnkProgram struct {
	Components []struct {
		Type      int `json:"type"`
		Component struct {
			Schedule []ProgramItem `json:"schedule"`
		} `json:"component"`
	} `json:"components"`
}

type ProgramItem struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	StartTime    string `json:"startTime"`
	LiveUrl      string `json:"liveUrl"`
	RecordingUrl string `json:"recordingUrl"`
	PosterImage  string `json:"posterImage"`
	Progress     int    `json:"progress"`
}
