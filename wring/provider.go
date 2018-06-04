package wring

import (
	"github.com/deadcheat/gong/interfaces"
	"github.com/deadcheat/gong/types"
	"github.com/deadcheat/gong/values"
)

type Provider struct{}

func New() interfaces.Provider {
	return &Provider{}
}

func (p *Provider) Provide(req types.AlexaRequest) (interfaces.Handler, error) {
	switch req.Request.Type {
	case values.TypeLaunchRequest:
		return &LaunchHandler{}, nil
	case values.TypeIntentRequest:
		return &IntentHandler{}, nil
	case values.TypeSessionEndedRequest:
		return &SessionEndedHandler{}, nil
	}
	return &UnknownHandler{}, nil
}
