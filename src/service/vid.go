package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

// 设置不重定向
func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 1 {
		return errors.New("stopped after 1 redirects")
	}
	return nil
}

// 获取视频id
func GetVideoId(douyinUrl string) (string, error) {
	// 设置不重定向
	client := &http.Client{
		CheckRedirect: defaultCheckRedirect,
	}
	request, err := http.NewRequest("GET", douyinUrl, nil)
	if err != nil {
		return "", err
	}
	// 添加代理
	request.Header.Set("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/72.0.3626.101 Mobile/15E148 Safari/605.1")
	response, _ := client.Do(request)
	// 获取重定向链接
	respUrl, err := response.Location()
	if err != nil {
		return "", err
	}
	locationUrl := respUrl.String()
	// 获取视频id
	compileRegex := regexp.MustCompile("video/(.*)/")
	matchArr := compileRegex.FindStringSubmatch(locationUrl)
	if len(matchArr) == 0 {
		return "", errors.New("can not find video id")
	}
	// 返回视频id
	return matchArr[len(matchArr)-1], nil
}

type XBogusRsp struct {
	XBogus string `json:"X-Bogus"`
	Param  string `json:"param"`
}

func GetXBogus(videoId string) (XBogusRsp, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://www.douyin.com/aweme/v1/web/aweme/detail/?aweme_id=%s&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333", videoId)
	params := make(map[string]string, 2)
	params["url"] = url
	params["user_agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	bytesData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", "https://tiktok.iculture.cc/X-Bogus", bytes.NewBuffer([]byte(bytesData)))
	var xB XBogusRsp
	if err != nil {
		return xB, err
	}
	request.Header.Set("User-Agent", "FancyPig")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Host", "tiktok.iculture.cc")
	request.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		fmt.Println(response)
		return xB, errors.New("请求失败")
	}
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &xB)
	return xB, err
}

// 创建msToken，107位
func CreateMsToken(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
