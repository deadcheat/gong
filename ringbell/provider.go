package ringbell

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
		i := req.Intent
		if i != nil {
			switch i.Name {
			case values.IntentGongRing:
				return &GongRingIntentHandler{}, nil
			case values.IntentAMAZONHelpIntent:
				return &HelpIntentHandler{}, nil
			case values.IntentAMAZONCancelIntent, values.IntentAMAZONStopIntent:
				return &SessionEndedHandler{}, nil
			}
		}
	case values.TypeSessionEndedRequest:
		return &SessionEndedHandler{}, nil
	}
	return &UnknownHandler{}, nil
}
