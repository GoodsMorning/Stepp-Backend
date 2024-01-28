package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	GetHelloHandler(c *gin.Context)
}