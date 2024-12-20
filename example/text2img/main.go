package main

import (
	"fmt"
	"github.com/alphasnow/liblib"
	"github.com/alphasnow/liblib/consts"
	"log"
)

func main() {
	api := liblib.NewAPI("AccessKey", "SecretKey")
	req := []byte(`{
    "templateUuid": "6f7c4652458d4802969f8d089cf5b91f",
    "generateParams": {
        "prompt": "filmfotos, Asian portrait,A young woman wearing a green baseball cap,covering one eye with her hand",
        "steps": 20,
        "width": 768,
        "height": 1024,
        "imgCount": 1, 
        "seed": -1,
        "restoreFaces": 0,
        "additionalNetwork": [
            {
                "modelId": "169505112cee468b95d5e4a5db0e5669",
                "weight": 1.0
            }
        ]
    }
}`)
	resp, err := api.Generate(consts.Text2Img, req)
	fmt.Println(resp, err)
	if err != nil {
		log.Fatal("发生错误:", err)
	}
	if resp.Code != 0 {
		log.Fatal("请求失败:", resp)
	}
	log.Println(resp.Data.GenerateUuid, err)
}
