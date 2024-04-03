package api

import (
	"github.com/Projects/channel-between-websocket-and-webhook-api/api/handler"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	corsConfig.AllowWebSockets = true
	r.Use(cors.New(corsConfig))

	h := handler.NewHandlerV1(&handler.HandlerV1Config{})
	r.GET("ws", h.GetLocation)
	r.POST("/zanjeer/devices", h.ForwardToWebhook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("swagger/doc.json")))

	return r
}
