package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deadcheat/gong/ringbell"
	"github.com/deadcheat/gong/types"
)

type MainHandler struct{}

func (h *MainHandler) Handle(req types.AlexaRequest) (types.AlexaResponse, error) {
	inner, _ := ringbell.New().Provide(req)
	return inner.Handle(req)
}

func main() {
	h := &MainHandler{}
	lambda.Start(h.Handle)
}
