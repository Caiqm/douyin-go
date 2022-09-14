package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Caiqm/douyin-go/src/model"
)

// 获取视频信息
func GetVideoApi(videoUrl string) (dyStruct model.DouYinVideo, err error) {
	videoId, err := GetVideoId(videoUrl)
	if err != nil {
		return
	}
	// 设置不重定向
	client := &http.Client{}
	douyinApiUrl := fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=%s", videoId)
	request, err := http.NewRequest("GET", douyinApiUrl, nil)
	if err != nil {
        return
    }
	// 添加代理
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		fmt.Println(response)
		err = errors.New("请求失败")
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
	_ = json.Unmarshal(body, &dyStruct)
	return
}