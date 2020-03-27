package initRouter

import (
	"io/ioutil"
	"sign/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()

	index := router.Group("/")
	{
		index.GET("", handler.Index)
	}

	return router
}
