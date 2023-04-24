package logs

import (
	"os"
	"time"

	"github.com/niubir/logger"
	"github.com/niubir/user-service/config"
	"github.com/niubir/utils"
)

var (
	logDir = "/logs"

	System *logger.Logger
)

func init() {
	if envLogDir := os.Getenv(`LOG_DIR`); envLogDir != "" {
		logDir = envLogDir
	}
	if !utils.FilepathExist(logDir) {
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			panic(err)
		}
	}
	utils.BoldPrint("***LOG DIR***\n%s", logDir)
}

func Init() error {
	var err error
	System, err = NewLogger("user_service")
	return err
}

func NewLogger(prefix string) (*logger.Logger, error) {
	logLevel := logger.ErrorLevel
	switch config.Config.LogLevel {
	case config.LOG_LEVEL_DEBUG:
		logLevel = logger.DebugLevel
	case config.LOG_LEVEL_INFO:
		logLevel = logger.InfoLevel
	case config.LOG_LEVEL_WARN:
		logLevel = logger.WarningLevel
	case config.LOG_LEVEL_ERROR:
		logLevel = logger.ErrorLevel
	}

	return logger.NewLogger(
		logger.WithTime(),
		logger.SetLevel(logLevel),
		logger.WithStack(),
		logger.WithStdout(),
		logger.WithFileout(
			logger.WithFilePath(logDir),
			logger.WithFilePrefix(prefix),
			logger.WithFileDuration(24*time.Hour),
			logger.WithFileMaxByte(10*1024*1024),
		),
	)
}
