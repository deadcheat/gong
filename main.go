package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deadcheat/gong/types"
	"github.com/deadcheat/gong/wring"
)

type MainHandler struct{}

func (h *MainHandler) Handle(req types.AlexaRequest) *types.AlexaResponse {
	inner, _ := wring.New().Provide(req)
	return inner.Handle(req)
}

func main() {
	h := &MainHandler{}
	lambda.Start(h)
}
