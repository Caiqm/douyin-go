package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 下载视频文件
func DownloadVideo(videoPlayUrl string) (path string, err error) {
	// 视频播放链接
	// videoPlayUrl := dyStruct.ItemList[0].Video.PlayAddr.UrlList[0]
	playUrl := strings.Replace(videoPlayUrl, "playwm", "play", -1)
	client := &http.Client{}
	request, _ := http.NewRequest("GET", playUrl, nil)
	request.Header.Set("user-agent", "Android")
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("请求失败")
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	dir, _ := os.Getwd()
	// 图片名称
	timeNano := time.Now().UnixNano()
	rand.Seed(timeNano)
	randNum := rand.Intn(9999)
	tmpName := timeNano + int64(randNum)
	filename := fmt.Sprintf("%s_%d.%s", "vid", tmpName, "mp4")
	path = filepath.Join(dir, "src", "video", filename)
	err = ioutil.WriteFile(path, body, 0777)
	return
}
