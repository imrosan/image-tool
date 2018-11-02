package main

import (
	"github.com/imrosan/image-tool/log"
	"github.com/imrosan/image-tool/server"
)

func main() {
	log.CreateLoggers()

	result := server.Start()
	if result < 0 {
		return
	}

	defer server.Stop()
}
