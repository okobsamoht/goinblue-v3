package goinblue

type SMS struct {
	To     string `json:"to"`
	Sender   string `json:"sender"`
	Text   string `json:"text"`
	WebUrl string `json:"web_url"`
	Tag    string `json:"tag"`
	Type   string `json:"type"`
}
