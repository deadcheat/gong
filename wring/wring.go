package wring

import (
	"fmt"
	"os"

	"github.com/deadcheat/gong/types"
	"github.com/deadcheat/gong/values"
)

type IntentHandler struct{}

func (h *IntentHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	url := os.Getenv(values.EnvKeyURL)
	return types.AlexaResponse{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: &types.OutputSpeech{
				Type: values.TypeSSML,
				SSML: fmt.Sprintf(`<speak><audio src="%s" /></speak>`, url),
			},
			ShouldEndSession: true,
		},
	}, nil
}

type LaunchHandler struct{}

func (h *LaunchHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	return types.AlexaResponse{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: &types.OutputSpeech{
				Type: values.TypePlainText,
				Text: "いつでもどうぞ",
			},
			Reprompt: &types.Reprompt{
				OutputSpeech: &types.OutputSpeech{
					Type: values.TypePlainText,
					Text: "ゴング、や、開始、などのように言っていただければゴングを鳴らします",
				},
			},
			ShouldEndSession: false,
		},
	}, nil
}

type SessionEndedHandler struct{}

func (h *SessionEndedHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	return types.AlexaResponse{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: &types.OutputSpeech{
				Type: values.TypePlainText,
				Text: "またどうぞ",
			},
			ShouldEndSession: true,
		},
	}, nil
}

type UnknownHandler struct{}

func (h *UnknownHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	return types.AlexaResponse{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: &types.OutputSpeech{
				Type: values.TypePlainText,
				Text: "申し訳ありません、聞き取れませんでした",
			},
			Reprompt: &types.Reprompt{
				OutputSpeech: &types.OutputSpeech{
					Type: values.TypePlainText,
					Text: "ゴング、や、開始、などのように言っていただければゴングを鳴らします",
				},
			},
			ShouldEndSession: false,
		},
	}, nil
}
