package channels

type Channel struct {
	Id          string `json:"id"`
	ContentType string `json:"contentType"`
	StreamUrl   string `json:"streamUrl"`
}
