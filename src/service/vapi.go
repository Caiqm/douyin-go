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
	xb, err := GetXBogus(videoId)
	if err != nil {
		return
	}
	// 设置不重定向
	client := &http.Client{}
	request, err := http.NewRequest("GET", xb.Param, nil)
	if err != nil {
		return
	}
	// 添加代理
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	request.Header.Set("Referer", "https://www.douyin.com/")
	request.Header.Set("Cookie", "msToken="+CreateMsToken(107)+";odin_tt=324fb4ea4a89c0c05827e18a1ed9cf9bf8a17f7705fcc793fec935b637867e2a5a9b8168c885554d029919117a18ba69; ttwid=1%7CWBuxH_bhbuTENNtACXoesI5QHV2Dt9-vkMGVHSRRbgY%7C1677118712%7C1d87ba1ea2cdf05d80204aea2e1036451dae638e7765b8a4d59d87fa05dd39ff; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWNsaWVudC1jc3IiOiItLS0tLUJFR0lOIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLVxyXG5NSUlCRFRDQnRRSUJBREFuTVFzd0NRWURWUVFHRXdKRFRqRVlNQllHQTFVRUF3d1BZbVJmZEdsamEyVjBYMmQxXHJcbllYSmtNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVKUDZzbjNLRlFBNUROSEcyK2F4bXAwNG5cclxud1hBSTZDU1IyZW1sVUE5QTZ4aGQzbVlPUlI4NVRLZ2tXd1FJSmp3Nyszdnc0Z2NNRG5iOTRoS3MvSjFJc3FBc1xyXG5NQ29HQ1NxR1NJYjNEUUVKRGpFZE1Cc3dHUVlEVlIwUkJCSXdFSUlPZDNkM0xtUnZkWGxwYmk1amIyMHdDZ1lJXHJcbktvWkl6ajBFQXdJRFJ3QXdSQUlnVmJkWTI0c0RYS0c0S2h3WlBmOHpxVDRBU0ROamNUb2FFRi9MQnd2QS8xSUNcclxuSURiVmZCUk1PQVB5cWJkcytld1QwSDZqdDg1czZZTVNVZEo5Z2dmOWlmeTBcclxuLS0tLS1FTkQgQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tXHJcbiJ9")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Host", "www.douyin.com")
	request.Header.Set("Connection", "keep-alive")
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
