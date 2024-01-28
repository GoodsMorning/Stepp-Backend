package service

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (s *service)GetHello(ctx context.Context) (gin.H, error) {
	return gin.H{
		"greeting_message": "Hi from Stepp-Backend!",
	},nil
}