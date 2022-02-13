package channels

import (
	"haservice/config"
	"io/ioutil"
	"net/http"

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

func GetProgram(channelConfig *config.ConfigChannel) (*string, error) {
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
