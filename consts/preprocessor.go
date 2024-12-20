//go:generate stringer -type=Preprocessor -linecomment -output preprocessor_string.go

package consts

// Preprocessor 定义了预处理器的枚举值
type Preprocessor int

// 预处理器枚举值常量定义
const (
	// 无处理器
	None Preprocessor = 0 // 0

	// Canny（硬边缘）
	Canny Preprocessor = 1

	// depth_midas
	DepthMidas Preprocessor = 2

	// depth_leres (LeRes 深度图估算)
	DepthLeres Preprocessor = 3

	// depth_leres++
	DepthLeresPlus Preprocessor = 4

	// hed
	HED Preprocessor = 5

	// hed_safe
	HEDSafe Preprocessor = 6

	// normal_map
	NormalMap Preprocessor = 9

	// mediapipe_face
	MediaPipeFace Preprocessor = 7

	// openpose (OpenPose 姿态)
	OpenPose Preprocessor = 10

	// openpose hand (OpenPose 姿态及手部)
	OpenPoseHand Preprocessor = 11

	// openpose face (OpenPose 姿态及脸部)
	OpenPoseFace Preprocessor = 12

	// openpose_faceonly (OpenPose 仅脸部)
	OpenPoseFaceOnly Preprocessor = 13

	// openpose_full (OpenPose 姿态、手部及脸部)
	OpenPoseFull Preprocessor = 14

	// mlsd (M-LSD 直线线条检测)
	MLSD Preprocessor = 8

	// pidinet
	PIDINET Preprocessor = 17

	// pidinet_safe
	PIDINETSafe Preprocessor = 18

	// softedge_teed
	SoftEdgeTeed Preprocessor = 58

	// softedge_anyline
	SoftEdgeAnyline Preprocessor = 65

	// scribble_pidinet(涂鸦- 手绘)
	ScribblePIDINET Preprocessor = 20

	// scribble_xdog (涂鸦- 强化边缘)
	ScribbleXdog Preprocessor = 21

	// scribble_hed(涂鸦 -合成)
	ScribbleHED Preprocessor = 22

	// lineart_realistic (写实线稿提取)
	LineartRealistic Preprocessor = 29

	// lineart standard (标准线稿提取 -白底黑线反色)
	LineartStandard Preprocessor = 32

	// lineart coarse (粗略线稿提取)
	LineartCoarse Preprocessor = 30

	// lineart_anime (动漫线稿提取)
	LineartAnime Preprocessor = 31

	// lineart_anime_denoise(动漫线稿提取-去噪)
	LineartAnimeDenoise Preprocessor = 36

	// depth_zoe (ZoE 深度图估算)
	DepthZoe Preprocessor = 25

	// depth_hand_refiner
	DepthHandRefiner Preprocessor = 57

	// depth_anything
	DepthAnything Preprocessor = 64

	// segmentation
	Segmentation Preprocessor = 23

	// oneformer_coco
	OneFormerCoco Preprocessor = 27

	// oneformer_ade20k
	OneFormerAde20k Preprocessor = 28

	// anime_face_segment
	AnimeFaceSegment Preprocessor = 54

	// normal bae (Bae 法线贴图提取)
	NormalBae Preprocessor = 26

	// densepose
	DensePose Preprocessor = 55

	// densepose_parula
	DensePoseParula Preprocessor = 56

	// tile_resample(分块重采样)
	TileResample Preprocessor = 34

	// tile_colorfix
	TileColorfix Preprocessor = 43

	// tile_colorfix+sharp
	TileColorfixSharp Preprocessor = 44

	// blur_gaussian
	BlurGaussian Preprocessor = 52

	// reference_only
	ReferenceOnly Preprocessor = 37

	// reference_adain
	ReferenceAdain Preprocessor = 38

	// reference_adain+attn
	ReferenceAdainAttn Preprocessor = 39

	// ip-adapter_clip_sd15
	IPAdapterClipSd15 Preprocessor = 48

	// ip-adapter_clip_sdxl
	IPAdapterClipSdxl Preprocessor = 49

	// ip-adapter_clip_sdxl_plus_vith
	IPAdapterClipSdxlPlusVith Preprocessor = 61

	// clip_vision
	ClipVision Preprocessor = 15

	// color
	Color Preprocessor = 16

	// pidinet_sketch
	PIDINETSketch Preprocessor = 19

	// shuffle (随机洗牌)
	Shuffle Preprocessor = 33

	// recolor_luminance
	RecolorLuminance Preprocessor = 50

	// recolor_intensity
	RecolorIntensity Preprocessor = 51

	// inpaint_global_harmonious
	InpaintGlobalHarmonious Preprocessor = 40

	// inpaint_only
	InpaintOnly Preprocessor = 41

	// inpaint_only+lama
	InpaintOnlyLama Preprocessor = 42

	// ip-adapter_face_id
	IPAdapterFaceId Preprocessor = 62

	// ip-adapter_face_id_plus
	IPAdapterFaceIdPlus Preprocessor = 63

	// instant_id_face_keypoints
	InstantIDFaceKeypoints Preprocessor = 59

	// instant_id_face_embedding
	InstantIDFaceEmbedding Preprocessor = 60

	// invert (白底黑线反色)
	Invert Preprocessor = 35
)
