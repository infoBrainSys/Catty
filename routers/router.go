package routers

import (
	"GPT/controllers"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	app := gin.Default()

	app.GET("/", controllers.Index)

	return app
}
