package app

import (
	"cgw/pkg/api"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.StaticFile("/", "./web/index.html")
	router.StaticFile("/styles.css", "./web/styles.css")
	router.LoadHTMLGlob("./templates/*")

	api.BindRoutes(router)

	return router
}
