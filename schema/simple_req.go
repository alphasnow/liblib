package schema

// SimpleImageSize 用于存储生成图像的具体宽度和高度信息。
type SimpleImageSize struct {
	// 图像的宽度，取值范围为512~2048，用于明确生成图像的具体宽度尺寸。
	Width int `json:"width"`
	// 图像的高度，取值范围为512~2048，用于确定生成图像的具体高度尺寸。
	Height int `json:"height"`
}

// SimpleControlnet 用于存储与controlnet相关的高级设置参数。
type SimpleControlnet struct {
	// controlnet的控制类型，可选择不同的值来实现不同的构图控制效果，如line（线稿轮廓）、depth（空间关系）、pose（人物姿态）、IPAdapter（风格迁移）等，此为非必填项，可根据具体需求选择合适的值。
	ControlType string `json:"controlType"`
	// controlnet所使用的控制图像的地址，需是可公网访问的完整URL，用于为controlnet提供参考图像，此为非必填项，当需要使用controlnet功能且选择了相应控制类型时需提供该地址。
	ControlImage string `json:"controlImage"`
}
