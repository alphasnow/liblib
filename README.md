# liblib
LiblibAI-哩布哩布AI的API接口 https://www.liblib.art/

## 功能
- Text2Img
- Img2Img
- Text2ImgUltra
- Img2ImgUltra
- Text2ImgClassic
- Img2ImgClassic
- Status

## 示例
### 建立任务
```go
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
resp, err := api.Generate(liblib.Text2Img, req)
fmt.Println(resp, err)
// resp.Data.GenerateUuid
```

### 查询任务结果
```go
api := liblib.NewAPI("AccessKey", "SecretKey")
req := &schema.StatusReq{GenerateUuid: "GenerateUuid"}
resp, err := api.Status(req)
// resp.Data.GenerateStatus
```

## 开源协议
- MIT

## 相关参考
- https://liblibai.feishu.cn/wiki/UAMVw67NcifQHukf8fpccgS5n6d
- https://www.liblib.art/