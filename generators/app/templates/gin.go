package gin

import (
  "<%= myAppPath %>/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
  // initialize routes
	router = InitRoutes()
}

//Run starts a Gin server.
func Run() {
	router.Run(":" + strconv.Itoa(config.ServerPort))
}

//ReturnRouter returns a pointer to the engine
func GetRouter() *gin.Engine {
	return router
}

//SetTestMode sets the gin mode as test.
func SetTestMode() {
	gin.SetMode(gin.TestMode)
}

