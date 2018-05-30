package types

type AlexaResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`
	Response          `json:"response"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Reprompt struct {
	OutputSpeech `json:"outputSpeech"`
}

type Directive struct {
	Type         string `json:"type"`
	PlayBehavior string `json:"playBehavior"`
	AudioItem    `json:"audioItem"`
}

type Stream struct {
	Token                string `json:"token"`
	URL                  string `json:"url"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}

type AudioItem struct {
	Stream `json:"stream"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	OutputSpeech     `json:"outputSpeech"`
	Card             `json:"card"`
	Reprompt         `json:"reprompt"`
	Directives       []Directive `json:"directives"`
	ShouldEndSession bool        `json:"shouldEndSession"`
}
