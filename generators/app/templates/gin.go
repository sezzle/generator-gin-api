package gin

import (
  "<%= myrepoUrl %>/<%= myappName %>/config"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	router *gin.Engine
	cfg    configE
)

//JSONError : Error for JSON return to frontend
type JSONError struct {
	Error string `json:"error"`
}

//config holds all of the environment variables for mysql db configuration.
type configE struct {
	Port    string `env:"GIN_PORT" envDefault:"10000"`
	Origins string `env:"GIN_CORS" envDefault:"http:localhost:4200"`
}

//Run starts a Gin server.
func Run() {
	router.Run(":" + strconv.Itoa(config.ServerPort))
}

//ReturnRouter returns a pointer to the engine
func ReturnRouter() *gin.Engine {
	return router
}

//SetTestMode sets the gin mode as test.
func SetTestMode() {
	gin.SetMode(gin.TestMode)
}

func init() {
	cfg = configE{}
	err := env.Parse(&cfg)
	if err != nil {
		glog.Info(err)
	}
	router = InitRoutes()
}
