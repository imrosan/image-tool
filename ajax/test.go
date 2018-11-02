package ajax

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/imrosan/image-tool/logger"
)

var log = logger.CreateLogger("test")

func Test(context *gin.Context) {
	log.Info("get a request")

    context.JSON(http.StatusOK, gin.H{
        "ret": 0,
        "msg": "hello, ok",
    })
}
