package schema

// Text2ImgClassicReq 用于存储生成图像任务的整体参数配置，包含参数模板的唯一标识以及生成图像的详细参数。
type Text2ImgClassicReq struct {
	// 参数模板的唯一标识，用于区分不同的参数模板设置，从预定义的模板集合中选取相应的值。
	TemplateUuid string `json:"templateUuid"`
	// 生成图像的详细参数结构体，涵盖了从图像特征描述、尺寸设置到高级控制等各方面的具体参数。
	GenerateParams Text2ImgClassicGenerateParams `json:"generateParams"`
}

// Text2ImgClassicGenerateParams 包含了生成图像过程中的各类详细参数，如提示词、图像尺寸、生成数量以及可选的高级设置等。
type Text2ImgClassicGenerateParams struct {
	// 正向提示词，用于描述期望生成图像的主要特征、风格、细节等内容，需为纯英文文本且不超过2000字符。
	Prompt string `json:"prompt"`
	// 图像宽高比预设值，可选择不同的预设比例来确定图像大致的宽高关系，与imageSize只能二选一进行配置。
	// 这里设置为"portrait"，表示宽高比为3:4，适合人物肖像，具体尺寸可参考如768*1024这种形式。
	AspectRatio string `json:"aspectRatio"`
	// 图像具体尺寸结构体，用于精确设置生成图像的宽度和高度，与aspectRatio只能二选一进行配置。
	ImageSize SimpleImageSize `json:"imageSize"`
	// 单次生成图像的数量，取值范围在1到4之间，明确一次生成任务要生成的图像张数。
	ImgCount int `json:"imgCount"`
	// controlnet相关的高级设置结构体，包含控制类型和控制图像等参数，用于实现更精细的图像生成控制，此部分可不填写。
	Controlnet SimpleControlnet `json:"controlnet"`
}
