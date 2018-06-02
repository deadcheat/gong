package wring

import (
	"github.com/deadcheat/gong/interfaces"
	"github.com/deadcheat/gong/types"
)

type Provider struct{}

func New() interfaces.Provider {
	return &Provider{}
}

func (p *Provider) Provide(req types.AlexaRequest) (interfaces.Handler, error) {
	return nil, nil
}
