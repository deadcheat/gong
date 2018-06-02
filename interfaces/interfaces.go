package interfaces

import "github.com/deadcheat/gong/types"

type Handler interface {
	Handle(req types.AlexaRequest) types.AlexaResponse
}

type Provider interface {
	Provide(req types.AlexaRequest) (Handler, error)
}
