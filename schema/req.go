package schema

// AdditionalNetwork 表示添加的LoRA相关信息
type AdditionalNetwork struct {
	ModelId string  `json:"modelId,omitempty"` // LoRA的模型版本uuid,从提供的LoRA列表中选择,非必填
	Weight  float64 `json:"weight,omitempty"`  // LoRA权重,-4.00 ~ +4.00,默认0.8,非必填
}

// InpaintParam 蒙版重绘相关参数
type InpaintParam struct {
	// 蒙版文件地址,只用文件名png,要求是白色蒙版、黑色底色的蒙版图URL,生图模式为蒙版重绘（mode=4）时必填
	MaskImage string `json:"maskImage,omitempty"`
	// 蒙版模糊度,取值范围是0 ~ 64,默认值为4,生图模式为蒙版重绘（mode=4）时必填
	MaskBlur int `json:"maskBlur,omitempty"`
	// 蒙版边缘预留像素,也称蒙版扩展量,取值范围是0 ~ 256,默认值为32,生图模式为蒙版重绘（mode=4）时必填
	MaskPadding int `json:"maskPadding,omitempty"`
	// 蒙版模式,0表示Inpaint_masked（重绘蒙版区域）,1表示Inpaint_not_masked（重绘非蒙版区域）,生图模式为蒙版重绘（mode=4）时必填
	MaskMode int `json:"maskMode,omitempty"`
	// 重绘区域,0表示whole_picture（重绘全图）,1表示only_masked（仅重绘蒙版区域）,生图模式为蒙版重绘（mode=4）时必填
	InpaintArea int `json:"inpaintArea,omitempty"`
	// 蒙版内容的填充模式,0表示fill（填充）,1表示original（原图）,2表示latent_noise（潜空间噪声）,3表示latent_nothing（空白潜空间）,生图模式为蒙版重绘（mode=4）时必填
	InpaintingFill int `json:"inpaintingFill,omitempty"`
}

// ControlNet 用于存储与ControlNet相关的各种参数设置
type ControlNet struct {
	// Controlnet单元顺序,取值范围为1 ~ 4,必填项
	UnitOrder int `json:"unitOrder"`
	// 图片地址,需是可公网访问的完整url,必填项
	SourceImage string `json:"sourceImage"`
	// 参考图宽度,不超过4096,必填项
	Width int `json:"width"`
	// 参考图高度,不超过4096,必填项
	Height int `json:"height"`
	// 预处理器枚举值,需从Controlnet预处理器列表中选择,必填项
	Preprocessor int `json:"preprocessor"`
	// 预处理参数,具体配置需参考预处理器参数配置,必填项
	AnnotationParameters map[string]interface{} `json:"annotationParameters"`
	// Controlnet模型uuid,需从提供的controlnet模型列表中选择,必填项
	Model string `json:"model"`
	// controlnet权重,取值范围为0 ~ 2,默认值为1,必填项
	ControlWeight float64 `json:"controlWeight"`
	// controlnet生效起始step,输入的值实际是表示占采样步数的百分比,取值范围为0 ~ 1,默认值为0,必填项
	StartingControlStep float64 `json:"startingControlStep"`
	// controlnet生效终止step,输入的值实际是表示占采样步数的百分比,取值范围为0 ~ 1,默认值为1,必填项
	EndingControlStep float64 `json:"endingControlStep"`
	// 完美像素模式,0表示关闭,1表示开启,默认值为1,必填项
	PixelPerfect int `json:"pixelPerfect"`
	// 控制模式,0表示balanced（均衡）,1表示prompt_important（更注重提示词）,2表示controlnet_important（更注重controlnet）,必填项
	ControlMode int `json:"controlMode"`
	// 缩放模式,0表示just_resize（直接缩放）,1表示crop_and_resize（裁剪并缩放）,2表示resize_and_fill（缩放并填充）,必填项
	ResizeMode int `json:"resizeMode"`
	// mask图片地址,蒙版图url,务必与参考图尺寸一致,要求是白色蒙版、黑色底色,非必填项
	MaskImage string `json:"maskImage,omitempty"`
}

// HiResFixInfo 用于存储高分辨率修复相关的各项参数信息。
type HiResFixInfo struct {
	// 高清修复采样步数，用于控制在高分辨率修复过程中进行采样的次数，取值范围为1到30，该参数为非必填项。
	HiresSteps int `json:"hiresSteps,omitempty"`
	// 高清修复去噪强度，用于确定在高分辨率修复时去除噪声的程度，取值范围是0到1且精确到百分位，此参数为非必填项。
	HiresDenoisingStrength float64 `json:"hiresDenoisingStrength,omitempty"`
	// 放大算法枚举值，需从提供的放大算法模型枚举中选择，用于指定在高分辨率修复中采用的放大算法模型，该参数为非必填项。
	Upscaler int `json:"upscaler,omitempty"`
	// 缩放宽度，明确在高分辨率修复或相关操作后图像的宽度尺寸，取值范围为128到2048，该参数为非必填项。
	ResizedWidth int `json:"resizedWidth,omitempty"`
	// 缩放高度，确定在高分辨率修复或相关操作后图像的高度尺寸，取值范围为128到2048，该参数为非必填项。
	ResizedHeight int `json:"resizedHeight,omitempty"`
}
