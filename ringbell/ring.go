package ringbell

import (
	"fmt"
	"os"

	"github.com/deadcheat/gong/types"
	"github.com/deadcheat/gong/values"
)

type GongRingIntentHandler struct{}

func (h *GongRingIntentHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	url := os.Getenv(values.EnvKeyURL)
	return types.AlexaResponse{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: &types.OutputSpeech{
				Type: values.TypeSSML,
				SSML: fmt.Sprintf(`<speak><prosody volume="x-loud"><audio src="%s" /></prosody></speak>`, url),
			},
			ShouldEndSession: true,
		},
	}, nil
}

type HelpIntentHandler struct{}

func (h *HelpIntentHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	return types.AlexaResponse{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: &types.OutputSpeech{
				Type: values.TypePlainText,
				Text: "このスキルは、合図を元にプロレスのゴングを鳴らすことができます。ゴングを鳴らすには、ゴング、や、開始、などのように合図してみてください。",
			},
			Reprompt: &types.Reprompt{
				OutputSpeech: &types.OutputSpeech{
					Type: values.TypePlainText,
					Text: "このスキルは、合図を元にプロレスのゴングを鳴らすことができます。ゴングを鳴らすには、ゴング、や、開始、などのように合図してみてください。",
				},
			},
			ShouldEndSession: false,
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
					Text: "ゴングを鳴らすには、ゴング、や、開始、などのように合図してください。",
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
				Text: "またゴングを鳴らしたくなったら呼んでください",
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
				Text: "申し訳ありません、聞き取れませんでした。ゴング、や、開始、などのように話しかけてみてください。",
			},
			Reprompt: &types.Reprompt{
				OutputSpeech: &types.OutputSpeech{
					Type: values.TypePlainText,
					Text: "ゴングを鳴らすには、ゴング、や、開始、などのように話しかけてみてください。",
				},
			},
			ShouldEndSession: false,
		},
	}, nil
}
