package channels

import (
	"haservice/config"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func GetStreamUrl(channelConfig *config.ConfigChannel) (*string, error) {
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
