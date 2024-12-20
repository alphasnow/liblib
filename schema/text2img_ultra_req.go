package schema

// Text2ImgUltraReq 用于存储生成图像任务的整体参数配置，包含参数模板ID以及生成图像的详细参数。
type Text2ImgUltraReq struct {
	// 参数模板的唯一标识，用于区分不同的参数模板设置。
	TemplateUuid string `json:"templateUuid"`
	// 生成图像的详细参数结构体，包含了从模型选择到图像生成各个环节的具体设置。
	GenerateParams Text2ImgUltraGenerateParams `json:"generateParams"`
}

// Text2ImgUltraGenerateParams 涵盖了生成图像过程中涉及的各类详细参数，如提示词、图像尺寸、图片数量等，以及高级设置（可选填）。
type Text2ImgUltraGenerateParams struct {
	// 正向提示词，以纯英文文本形式描述期望生成图像的特征、内容等，字符长度不超过2000字符，此为必填项，对生成图像的最终效果起到引导作用。
	Prompt string `json:"prompt"`
	// 图片宽高比预设，可选择不同的预设值来确定图像的大致宽高比例或具体尺寸，与imageSize二者只能选其一进行配置。
	// 可选值如square（宽高比1:1，通用，具体尺寸1024*1024）、portrait（宽高比3:4，适合人物肖像，具体尺寸768*1024）、landscape（宽高比16:9，适合影视画幅，具体尺寸1280*720）等。
	AspectRatio string `json:"aspectRatio"`
	// 图片具体宽高设置结构体，用于设置生成图像的具体宽度和高度，与aspectRatio二者只能选其一进行配置，宽度取值范围为512~2048，高度取值范围为512~2048。
	ImageSize SimpleImageSize `json:"imageSize"`
	// 单次生成图像的数量，取值范围为1~4，此为必填项，明确一次生成任务要产出的图像张数。
	ImgCount int `json:"imgCount"`
	// 构图控制相关参数结构体，用于设置如controlnet等高级功能的相关参数，此部分为非必填项，可根据需求进行配置。
	Controlnet SimpleControlnet `json:"controlnet"`
}
