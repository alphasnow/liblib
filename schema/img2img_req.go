package schema

// Img2ImgReq 包含了生成图像所需的各种参数设置
type Img2ImgReq struct {
	// 预设模版uuid,从提供的预设参数模版中选择
	TemplateUuid string `json:"templateUuid"`
	// 生成图像的具体参数配置
	GenerateParams Img2ImgGenerateParams `json:"generateParams"`
}

// Img2ImgGenerateParams 图生图基础参数
type Img2ImgGenerateParams struct {
	// 模型的唯一标识,需从提供的Checkpoint列表里选择,此为必填项,用于指定生成图像所依据的基础模型。
	CheckPointId string `json:"checkPointId"`
	// LoRA模型的附加组合及各自参数的列表,可根据具体需求配置多个LoRA模型来影响生成图像的效果,非必填项,其具体参数配置可参考相关说明。
	AdditionalNetwork []AdditionalNetwork `json:"additionalNetwork,omitempty"`
	// VAE的模型唯一标识,若需要可从提供的VAE列表中选取,非必填项,用于特定的图像生成处理方式（如变分自编码器相关操作）。
	VaeId string `json:"vaeId,omitempty"`
	// 正向提示词,以纯英文文本形式描述期望生成图像的特征、内容等,字符长度不超过2000字符,此为必填项,对生成图像的最终效果起到引导作用。
	Prompt string `json:"prompt"`
	// 负向提示词,同样为纯英文文本,用于明确不希望出现在生成图像中的特征、情况等,字符长度不超过2000字符,是必填项,可辅助排除不符合要求的图像特征。
	NegativePrompt string `json:"negativePrompt"`
	// Clip跳过层的设置值,决定在生成图像过程中跳过的特定层数,取值范围在1到12之间,此为必填项,用于调整生成图像的某些内部处理逻辑。
	ClipSkip int `json:"clipSkip"`
	// 采样器枚举值,需从预定义的采样方法列表中挑选合适的值,用于确定生成图像时采用的具体采样方式,此为必填项。
	Sampler int `json:"sampler"`
	// 采样步数,即生成图像过程中进行采样操作的次数,取值范围为1到60,此为必填项,影响生成图像的精细程度和质量。
	Steps int `json:"steps"`
	// cfg_scale的值,用于控制提示词对生成图像的影响程度,取值范围在1.0到15.0之间,此为必填项。
	CfgScale float64 `json:"cfgScale"`
	// 随机种子来源类型,0表示CPU,1表示GPU,用于确定生成图像时随机数的产生来源,此为必填项,不同来源可能会对生成结果的随机性产生影响。
	RandnSource int `json:"randnSource"`
	// 随机种子值,其范围从 -1到9999999999,当值为 -1时表示随机生成种子,该参数是必填项,相同种子在相同条件下可复现生成图像的结果。
	Seed int `json:"seed"`
	// 单次生成图像的数量,取值范围为1到4,此为必填项,明确一次生成任务要产出的图像张数。
	ImgCount int `json:"imgCount"`
	// 面部修复设置,0表示关闭面部修复功能,1表示开启,默认值为0,用于控制是否对生成图像中的面部进行特殊修复处理,此为必填项。
	RestoreFaces int `json:"restoreFaces"`
	// 参考图地址,需是可公网访问的完整URL,在某些生成模式（如图生图等）下可能会依据此参考图进行相关操作,此为必填项。
	SourceImage string `json:"sourceImage"`
	// 缩放模式的设置值,0表示just_resize（直接缩放）,1表示crop_and_resize（裁剪并缩放）,2表示resize_and_fill（缩放并填充）,此为必填项,用于确定对图像进行尺寸调整的具体方式。
	ResizeMode int `json:"resizeMode"`
	// 调整后的图片宽度,取值范围在128到2048之间,此为必填项,明确生成图像经过缩放等操作后最终的宽度尺寸。
	ResizedWidth int `json:"resizedWidth"`
	// 调整后的图片高度,取值范围同样在128到2048之间,此为必填项,确定生成图像经过缩放等操作后最终的高度尺寸。
	ResizedHeight int `json:"resizedHeight"`
	// 生图模式的设置值,0表示img2img（图生图）,4表示inpaint_upload_mask（蒙版重绘）,此为必填项,用于指定本次生成图像所采用的具体模式。
	Mode int `json:"Mode"`
	// 去噪强度（图生图重绘幅度）的值,取值范围在0到1之间,默认值为0.75,用于控制在图生图过程中对图像进行重绘时的去噪程度,此为必填项。
	DenoisingStrength float64 `json:"denoisingStrength"`
	// 蒙版重绘相关参数结构体,当生图模式为蒙版重绘（Mode = 4）时必填,其具体参数配置可参考相关说明,用于设置蒙版重绘操作的各项细节。
	InpaintParam InpaintParam `json:"inpaintParam,omitempty"`
	// 模型加载的ControlNet组合及各自参数的列表,可根据具体需求配置多个ControlNet来辅助生成图像,非必填项,其具体参数配置可参考相关说明。
	ControlNet []ControlNet `json:"ControlNet,omitempty"`
}
