package service

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetHello(ctx context.Context) (gin.H, error)
}