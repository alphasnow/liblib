package schema

// Text2ImgReq 包含了生成图像所需的各种参数设置
type Text2ImgReq struct {
	// 预设模版uuid,从提供的预设参数模版中选择
	TemplateUuid string `json:"templateUuid"`
	// 生成图像的具体参数配置
	GenerateParams Text2ImgGenerateParams `json:"generateParams"`
}

// Text2ImgGenerateParams 文生图基础参数
type Text2ImgGenerateParams struct {
	// 模型的唯一标识，需从提供的Checkpoint列表里选择，此为必填项，用于指定生成图像所依据的基础模型。
	CheckPointId string `json:"checkPointId"`
	// LoRA组合及权重设置的列表，可根据具体需求配置多个LoRA模型来影响生成图像的效果，非必填项，
	// 其具体参数配置可参考相关说明，且LoRA的基础算法类型需要与checkpoint一致。
	AdditionalNetwork []AdditionalNetwork `json:"additionalNetwork,omitempty"`
	// VAE的模型唯一标识，若需要可从提供的VAE列表中选取，非必填项，可为空，空值表示取checkpoint的VAE，
	// 用于特定的图像生成处理方式（如变分自编码器相关操作）。
	VaeId string `json:"vaeId,omitempty"`
	// 正向提示词，以纯英文文本形式描述期望生成图像的特征、内容等，字符长度不超过2000字符，此为必填项，
	// 对生成图像的最终效果起到引导作用。
	Prompt string `json:"prompt"`
	// 负向提示词，同样为纯英文文本，用于明确不希望出现在生成图像中的特征、情况等，字符长度不超过2000字符，
	// 是必填项，可辅助排除不符合要求的图像特征。
	NegativePrompt string `json:"negativePrompt"`
	// Clip跳过层的设置值，取值范围在1到12之间，默认值为2，此为必填项，用于调整生成图像的某些内部处理逻辑。
	ClipSkip int `json:"clipSkip"`
	// 采样器枚举值，需从预定义的采样方法列表中挑选合适的值，用于确定生成图像时采用的具体采样方式，此为必填项。
	Sampler int `json:"sampler"`
	// 采样步数，即生成图像过程中进行采样操作的次数，取值范围为1到60，此为必填项，影响生成图像的精细程度和质量。
	Steps int `json:"steps"`
	// cfg_scale的值，用于控制提示词对生成图像的影响程度，取值范围在1.0到15.0之间，此为必填项。
	CfgScale float64 `json:"cfgScale"`
	// 初始宽度，取值范围在128到1536之间，不同基础算法有不同的建议区间，如基础算法1.5建议区间为512~768，
	// 基础算法XL建议区间为768~1344，基础算法F.1建议区间为768~2048，此为必填项，明确生成图像的初始宽度。
	Width int `json:"width"`
	// 初始高度，取值范围在128到1536之间，不同基础算法有不同的建议区间，如基础算法1.5建议区间为512~768，
	// 基础算法XL建议区间为768~1344，基础算法F.1建议区间为768~2048，此为必填项，确定生成图像的初始高度。
	Height int `json:"height"`
	// 单次生成图像的数量，取值范围为1到4，此为必填项，明确一次生成任务要产出的图像张数。
	ImgCount int `json:"imgCount"`
	// 随机种子生成来源，0表示CPU，1表示GPU，默认值为0，此为必填项，用于确定生成图像时随机数的产生来源，
	// 不同来源可能会对生成结果的随机性产生影响。
	RandnSource int `json:"randnSource"`
	// 随机种子值，其范围从 -1到9999999999，当值为 -1时表示随机生成种子，该参数是必填项，相同种子在相同条件下可复现生成图像的结果。
	Seed int64 `json:"seed"`
	// 面部修复设置，0表示关闭面部修复功能，1表示开启，默认值为0，用于控制是否对生成图像中的面部进行特殊修复处理，此为必填项。
	RestoreFaces int `json:"restoreFaces"`
	// 高分辨率修复相关参数结构体，非必填项，其具体参数配置可参考高分辨率修复的相关参数，用于设置高分辨率修复的各项细节。
	HiResFixInfo HiResFixInfo `json:"hiResFixInfo,omitempty"`
	// 模型加载的ControlNet组合及各自参数的列表，可根据具体需求配置多个ControlNet来辅助生成图像，非必填项，
	// 其具体参数配置可参考controlnet参数配置。
	ControlNet []ControlNet `json:"controlNet,omitempty"`
}
