package main

import (
	"github.com/alphasnow/liblib"
	"github.com/alphasnow/liblib/consts"
	"github.com/alphasnow/liblib/schema"
	"log"
	"time"
)

func main() {
	api := liblib.NewAPI("AccessKey", "SecretKey")
	req := &schema.StatusReq{GenerateUuid: "GenerateUuid"}

	timeout := 600 * time.Second
	start := time.Now()
	for {
		now := time.Now()
		if now.Sub(start) >= timeout {
			log.Println("请求超时")
			break
		}
		resp, err := api.Status(req)
		if err != nil {
			log.Println("请求失败:", err)
			break
		}

		if resp.Code != 0 {
			log.Println("请求失败:", resp)
			break
		}

		if resp.Data.GenerateStatus == int(consts.Failure) {
			log.Println("请求失败:", resp)
			break
		}
		if resp.Data.GenerateStatus == int(consts.Success) {
			for _, img := range resp.Data.Images {
				log.Println("请求成功:", img.ImageUrl)
			}
			break
		}

		log.Println("5秒后重新请求")
		time.Sleep(5 * time.Second)
	}
}
