package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag"
	"github.com/zchat-team/zchat-server/swagger"
)

func Swagger(r gin.IRouter) {
	r.GET("/swagger/*any", SwaggerHandler())
}

func SwaggerHandler() gin.HandlerFunc {
	swag.Register(swag.Name, new(swagger.Docs))
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
