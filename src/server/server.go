package server

import (
	"context"
	"stepp-backend/initializer"
	"stepp-backend/src/handler"
	"stepp-backend/src/router"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.InitEnv()
	initializer.ConnectDB()
}

func RunServer() error {
	r := gin.Default()
	ctx := context.Background()

	serverHandler := handler.NewHandler(ctx)
	router.Router(r, serverHandler)
	
	return r.Run()
}
