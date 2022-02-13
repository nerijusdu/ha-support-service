package channels

import (
	"fmt"
	"haservice/config"
	"haservice/utils"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

func GetStreamUrl(channelConfig *config.ConfigChannel) (*string, error) {
	res, err := http.Get(channelConfig.Url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	strBody := string(body)
	streamUrl := gjson.Get(strBody, channelConfig.StreamUrlPath).String()

	return &streamUrl, nil
}

func GetProgram(w http.ResponseWriter, channelConfig *config.ConfigChannel) {
	var content template.HTML
	var data []ProgramItem
	var err error

	if channelConfig.Id == "lnk" {
		data, err = getProgramFromJson(channelConfig)
	} else {
		var program *string
		program, err = getProgramWithSelector(channelConfig)
		content = template.HTML(*program)
	}

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hostUrl := os.Getenv("HOST_URL")
	tmpl := template.Must(template.ParseFiles("templates/channels/program.html"))
	tmpl.Execute(w, ProgramTemplateData{
		Host:       hostUrl,
		Id:         channelConfig.Id,
		Content:    content,
		Data:       data,
		Stylesheet: channelConfig.ProgramStylesheet,
		ScrollTo:   channelConfig.ProgramScrollTo,
	})
}

func getProgramWithSelector(channelConfig *config.ConfigChannel) (*string, error) {
	res, err := http.Get(channelConfig.ProgramUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	programHtml, err := goquery.OuterHtml(doc.Find(channelConfig.ProgramSelector).First())
	return &programHtml, err
}

func getProgramFromJson(channelConfig *config.ConfigChannel) ([]ProgramItem, error) {
	program := &LnkProgram{}
	results := []ProgramItem{}
	err := utils.GetJson(
		channelConfig.ProgramUrl,
		program,
	)
	if err != nil {
		return nil, err
	}

	for _, comp := range program.Components {
		if comp.Type == 19 {
			results = comp.Component.Schedule
			break
		}
	}

	for i := 0; i < len(results); i++ {
		t, _ := time.Parse("2006-01-02T15:04:05", results[i].StartTime)
		results[i].StartTime = t.Format("15:04")
		results[i].PosterImage = "https://lnk.lt/all-images/" + results[i].PosterImage
	}

	return results, nil
}
