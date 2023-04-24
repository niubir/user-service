package http

import (
	"fmt"
	"io"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/niubir/user-service/config"
	"github.com/niubir/user-service/logs"
)

func Init(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	if !config.Config.HTTP.Enable {
		return
	}

	address := fmt.Sprintf(":%d", config.Config.HTTP.Port)

	setGinMode()
	setGinLogger()
	engine := gin.Default()
	initRoutes(engine)

	logs.System.Infof("start http, address(%s)\n", address)
	engine.Run(address)
}

func setGinLogger() {
	ginWriter := io.Discard
	ginErrorWriter := io.Discard
	if config.Config.HTTP.Debug {
		if w, err := logs.NewLogger("user_service_http"); err != nil {
			logs.System.Errorln("create user service http logger failed:", err)
		} else {
			ginWriter = w
		}
		if w, err := logs.NewLogger("user_service_http_error"); err != nil {
			logs.System.Errorln("create user service error http logger failed:", err)
		} else {
			ginErrorWriter = w
		}
	}
	gin.DefaultWriter = ginWriter
	gin.DefaultErrorWriter = ginErrorWriter
}

func setGinMode() {
	var ginMode string
	switch {
	case config.Config.IsProduction():
		ginMode = gin.ReleaseMode
	case config.Config.IsTest():
		ginMode = gin.TestMode
	default:
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
}
