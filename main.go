package main

import (
	"github.com/Caiqm/douyin-go/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	r := route.Group("api")
	{
		r.GET("video", api.GetVideoInfo)
		r.GET("download", api.DownloadVideo)
	}
	route.Run(":8097")
}