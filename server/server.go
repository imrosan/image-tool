package server 

import (
	"github.com/gin-gonic/gin"
	"github.com/facebookgo/grace/gracehttp"
	"go.uber.org/zap"
	"github.com/imrosan/image-tool/log"
	"github.com/imrosan/image-tool/ajax"
	"net/http"
	"time"
)

var server http.Server

func Start() int {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	var engine *gin.Engine = gin.New()
	var group *gin.RouterGroup = engine.Group("/ajax")
	group.GET("/test", ajax.Test)

	server = http.Server{
		Addr: "0.0.0.0:80",
		Handler: engine, 
		ReadTimeout: time.Duration(300) * time.Second,
		WriteTimeout: time.Duration(300) * time.Second,
	}

	log.Main.Info("start http server")
	err := gracehttp.Serve(&server)
	if err != nil {
		log.Main.Error("http server start failed",
			zap.String("error", err.Error()),
		)

		return -1
	}

	return 0
}

func Stop() {
	err := server.Close()
	if err != nil {
		log.Main.Error("http server shut down failed",
			zap.String("error", err.Error()),
		)
	} else {
		log.Main.Info("http server shut down success")
	}
}
