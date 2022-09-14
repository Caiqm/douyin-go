package api

import (
	"net/http"

	"github.com/Caiqm/douyin-go/src/service"
	"github.com/gin-gonic/gin"
)

// 获取视频信息
func GetVideoInfo(ctx *gin.Context) {
	videoLink := ctx.Query("link")
	if videoLink == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": "请输入视频链接",
			"data": gin.H{},
		})
		return
	}
	dy, err := service.GetVideoApi(videoLink)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": err.Error(),
			"data": gin.H{},
		})
		return
	}
	if dy.StatusCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": "请求接口失败",
			"data": gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "ok",
		"data": dy.ItemList,
	})
}

// 下载文件
func DownloadVideo(ctx *gin.Context) {
	videoLink := ctx.Query("link")
	if videoLink == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": "请输入视频链接",
			"data": gin.H{},
		})
		return
	}
	filename, err := service.DownloadVideo(videoLink)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": err.Error(),
			"data": gin.H{},
		})
		return
	}
	ctx.File(filename)
}