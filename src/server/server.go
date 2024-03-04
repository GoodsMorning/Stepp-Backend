package server

import (
	"context"
	"stepp-backend/src/handler"
	"stepp-backend/src/initializer"
	"stepp-backend/src/router"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.InitEnv()
}

func RunServer() error {
	r := gin.Default()
	ctx := context.Background()

	serverHandler := handler.NewHandler(ctx)
	router.Router(r, serverHandler)

	return r.Run()
}
