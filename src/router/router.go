package router

import (
	"stepp-backend/src/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, handler handler.Handler) {

	r.GET("/", handler.GetHelloHandler)

	// test := r.Group("/test")
	// {
	// 	test.GET("/", )
	// }

}