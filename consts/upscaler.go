//go:generate stringer -type=Upscaler -linecomment -output upscaler_string.go

package consts

// Upscaler 放大算法模型
type Upscaler int

const (
	// Latent: 传统放大，低分辨率图像的适度放大，对原图保留度高
	Latent Upscaler = iota // 0

	// Latent (antialiased): 在Latent的基础上增加了抗锯齿处理，适合需要平滑边缘的图像
	LatentAntialiased Upscaler = iota // 1

	// Latent (bicubic): 使用双三次插值算法，适合需要较高质量放大的图像
	LatentBicubic Upscaler = iota // 2

	// Latent (bicubic antialiased): 结合双三次插值和抗锯齿，适合高质量且平滑的图像放大
	LatentBicubicAntialiased Upscaler = iota // 3

	// Latent (nearest): 使用最近邻插值，速度快但质量较低，适合简单图形
	LatentNearest Upscaler = iota // 4

	// Latent (nearest-exact): 使用精确的最近邻插值算法，适合需要保留原始像素的图像
	LatentNearestExact Upscaler = iota // 5

	// Lanczos: 一种高质量的插值算法，适合需要高保真度的图像放大
	Lanczos Upscaler = iota // 6

	// Nearest: 最近邻插值，简单快速，适合低质量图像放大
	Nearest Upscaler = iota // 7

	// ESRGAN_4x: 基于增强型超分辨率生成对抗网络，适合高质量图像放大
	ESRGAN_4x Upscaler = iota // 8

	// LDSR: 基于深度学习的超分辨率算法，适合处理复杂细节的图像
	LDSR Upscaler = iota // 9

	// R-ESRGAN_4x+: 改进版的ESRGAN，适合高质量放大
	R_ESRGAN_4xPlus Upscaler = iota // 10

	// R-ESRGAN_4x+ Anime6B: 针对动漫图像优化的R-ESRGAN
	R_ESRGAN_4xPlus_Anime6B Upscaler = iota // 11

	// ScuNET GAN: 基于生成对抗网络的超分辨率方法
	ScuNET_GAN Upscaler = iota // 12

	// ScuNET PSNR: 相较于ScuNET GAN，PSNR版本更注重图像质量
	ScuNET_PSNR Upscaler = iota // 13

	// SwinIR_4x: 基于Swin Transformer的超分辨率方法
	SwinIR_4x Upscaler = iota // 14

	// 4x-UltraSharp: 专注于图像锐化的超分辨率算法
	C4x_UltraSharp Upscaler = iota // 15

	// 8x-NMKD-Superscale: 专注于高倍放大的算法
	C8x_NMKD_Superscale Upscaler = iota // 16

	// 4x_NMKD-Siax_200k: 4倍放大算法，基于特定数据集训练
	C4x_NMKD_Siax_200k Upscaler = iota // 17

	// 4x_NMKD-Superscale-SP_178000_G: 4倍放大算法，基于不同数据集训练
	C4x_NMKD_Superscale_SP_178000_G Upscaler = iota // 18

	// 4x-AnimeSharp: 针对动漫图像的锐化和放大算法
	C4x_AnimeSharp Upscaler = iota // 19

	// 4x_foolhardy_Remacri: 强调细节恢复的放大算法
	C4x_Foolhardy_Remacri Upscaler = iota // 20

	// BSRGAN: 基于对抗学习的超分辨率方法
	BSRGAN Upscaler = iota // 21

	// DAT 2: 主要用于图像解压和放大
	DAT2 Upscaler = iota // 22

	// DAT 3: 主要用于图像解压和放大
	DAT3 Upscaler = iota // 23

	// DAT 4: 主要用于图像解压和放大
	DAT4 Upscaler = iota // 24

	// 4x-DeCompress: 主要用于图像解压和放大
	C4x_DeCompress Upscaler = iota // 25

	// 4x-DeCompress Strong: 主要用于图像解压和放大，增强材质效果
	C4x_DeCompress_Strong Upscaler = iota // 26
)
