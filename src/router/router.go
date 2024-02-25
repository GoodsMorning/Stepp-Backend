package router

import (
	"stepp-backend/src/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, handler handler.Handler) {
	r.GET("/", handler.GetHelloHandler)
	r.GET("/post/:post-id", handler.GetPostHandler)
}
