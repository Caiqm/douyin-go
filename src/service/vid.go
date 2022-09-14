package service

import (
	"errors"
	"net/http"
	"regexp"
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
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	response, _ := client.Do(request)
	// 获取重定向链接
	respUrl, err := response.Location()
	if err != nil{
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