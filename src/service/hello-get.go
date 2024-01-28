package service

import (
	"context"
	basepayload "stepp-backend/src/model/base"
)

func (s *service)GetHello(ctx context.Context) (*basepayload.BasePayload, error) {
	resp := struct {
		GreetingMessage string `json:"greeting_message"`
	}{
		GreetingMessage: "Hi from Stepp-Backend!",
	}
	return basepayload.GetSuccessResponse(resp),nil
}