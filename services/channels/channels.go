package channels

import (
	"fmt"
	"haservice/config"
	"haservice/utils"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

	if channelConfig.ProgramSelector == "" {
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
		Host:            hostUrl,
		Id:              channelConfig.Id,
		Content:         content,
		Data:            data,
		Stylesheet:      channelConfig.ProgramStylesheet,
		LocalStylesheet: channelConfig.ProgramLocalStylesheet,
		ScrollTo:        channelConfig.ProgramScrollTo,
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

	selector := channelConfig.ProgramSelector
	selector = strings.Replace(selector, "$date$", time.Now().Format("2006.01.02"), -1)

	programHtml, err := goquery.OuterHtml(doc.Find(selector).First())
	return &programHtml, err
}

func getProgramFromJson(channelConfig *config.ConfigChannel) ([]ProgramItem, error) {
	results := []ProgramItem{}
	err := utils.GetJson(
		fmt.Sprintf(channelConfig.ProgramUrl, time.Now().Format("2006-01-02T15:04:05.000Z")),
		&results,
	)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(results); i++ {
		t, _ := time.Parse("2006-01-02T15:04:05", results[i].StartTime)
		results[i].StartTime = t.Format("15:04")
		results[i].PosterImage = "https://lnk.lt/all-images/" + results[i].PosterImage
	}

	return results, nil
}
