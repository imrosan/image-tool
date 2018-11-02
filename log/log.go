package log

import (
	"go.uber.org/zap"
	"github.com/imrosan/image-tool/logger"
)

var Main *zap.Logger

func CreateLoggers() {
	Main = logger.CreateLogger("main")
}
