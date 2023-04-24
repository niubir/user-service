package http

import (
	"github.com/gin-gonic/gin"
	"github.com/niubir/user-service/http/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(engine *gin.Engine) {
	api := engine.Group("")

	docs.SwaggerInfo.Title = "user service"
	docs.SwaggerInfo.Version = "v0.0.1"

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
