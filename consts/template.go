package consts

// 进阶模式
const (
	// FlagshipF1TextToImage F.1 文生图模板 UUID
	// Checkpoint 默认为官方模型，可用模型范围为基础算法 F.1，支持 additional network
	FlagshipF1TextToImage = "6f7c4652458d4802969f8d089cf5b91f"

	// FlagshipF1ImageToImage F.1 图生图模板 UUID
	// Checkpoint 默认为官方模型，可用模型范围为基础算法 F.1，支持 additional network
	FlagshipF1ImageToImage = "63b72710c9574457ba303d9d9b8df8bd"

	// Classic15AndXLTextToImage 1.5 和 XL 文生图模板 UUID
	// 可用模型范围为基础算法 1.5 和基础算法 XL，支持 additional network、高分辨率修复和 controlnet
	// 可通过自由拼接参数实现各类的文生图诉求
	Classic15AndXLTextToImage = "e10adc3949ba59abbe56e057f20f883e"

	// Classic15AndXLImageToImage 1.5 和 XL 图生图模板 UUID
	// 可用模型范围为基础算法 1.5 和基础算法 XL，支持 additional network 和 controlnet
	// 可通过自由拼接参数实现各类的图生图和蒙版重绘诉求
	Classic15AndXLImageToImage = "9c7d531dc75f476aa833b3d452b8f7ad"

	// ControlnetLocalRedraw Controlnet局部重绘模板 UUID
	// 提取自文生图完整参数，支持 additional network 和 controlnet，不支持高分辨率修复（hiresfix）
	ControlnetLocalRedraw = "b689de89e8c9407a874acd415b3aa126"

	// ImageToImageLocalRedraw 图生图局部重绘模板 UUID
	// 提取自图生图完整参数，支持 additionalNetwork，不支持 Controlnet
	ImageToImageLocalRedraw = "74509e1b072a4c45a7f1843a963c8462"

	// InstantIDFaceSwap InstantID人像换脸模板 UUID
	// 仅用于人像换脸，注意人像参考图中的人物面部特征务必清晰
	InstantIDFaceSwap = "7d888009f81d4252a7c458c874cd017f"
)

// 简易模式
const (
	// SimpleFlagshipF1TextToImage 旗舰版F.1文生图 UUID
	// 旗舰版文生图模板的UUID，Checkpoint默认为官方旗舰版F.1模型，支持指定的几款Controlnet
	SimpleFlagshipF1TextToImage = "5d7e67009b344550bc1aa6ccbfa1d7f4"

	// SimpleFlagshipF1ImageToImage 旗舰版F.1图生图 UUID
	// 旗舰版图生图模板的UUID，Checkpoint默认为官方旗舰版F.1模型，支持指定的几款Controlnet
	SimpleFlagshipF1ImageToImage = "07e00af4fc464c7ab55ff906f8acf1b7"

	// SimpleClassicXLTextToImage 经典版XL文生图 UUID
	// 经典版文生图模板的UUID，Checkpoint默认为官方经典版XL模型，支持指定的几款Controlnet
	SimpleClassicXLTextToImage = "3af36dd5a61e4da88c6cb5eb57a8fe2e"

	// SimpleClassicXLImageToImage 经典版XL图生图 UUID
	// 经典版图生图模板的UUID，Checkpoint默认为官方经典版XL模型，支持指定的几款Controlnet
	SimpleClassicXLImageToImage = "e653a58128d34d1dbc231a03e4fedd6f"
)
