package wring

import "github.com/deadcheat/gong/types"

type IntentHandler struct{}

func (h *IntentHandler) Handle(req types.AlexaRequest) *types.AlexaResponse {
	return nil
}

type LaunchHandler struct{}

func (h *LaunchHandler) Handle(req types.AlexaRequest) *types.AlexaResponse {
	return nil
}

type SessionEndedHandler struct{}

func (h *SessionEndedHandler) Handle(req types.AlexaRequest) *types.AlexaResponse {
	return nil
}
