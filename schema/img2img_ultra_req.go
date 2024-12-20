package schema

// Img2ImgUltraReq 用于存储生成图像任务的整体参数配置，包含参数模板ID以及生成图像的详细参数。
type Img2ImgUltraReq struct {
	// 参数模板的唯一标识，用于区分不同的参数模板设置。
	TemplateUuid string `json:"templateUuid"`
	// 生成图像的详细参数结构体，包含了从模型选择到图像生成各个环节的具体设置。
	GenerateParams Img2ImgUltraGenerateParams `json:"generateParams"`
}

// Img2ImgUltraGenerateParams 涵盖了生成图像过程中涉及的各类详细参数，如提示词、图像尺寸、图片数量等，以及高级设置（可选填）。
type Img2ImgUltraGenerateParams struct {
	// 正向提示词，用于描述期望生成图像的主要特征、风格、细节等内容，通常为纯英文文本且不超过2000字符。
	Prompt string `json:"prompt"`
	// 参考图地址，需是可公网访问的完整URL，该参考图可作为生成图像的依据或参考，例如在图生图等模式下发挥作用。
	SourceImage string `json:"sourceImage"`
	// 单次生成图像的数量，取值范围在1到4之间，明确一次生成任务要生成的图像张数。
	ImgCount int `json:"imgCount"`
	// controlnet相关的高级设置结构体，包含控制类型和控制图像等参数，用于实现更精细的图像生成控制，此部分可不填写。
	Controlnet SimpleControlnet `json:"controlnet"`
}
