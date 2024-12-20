//go:generate stringer -type=Sampler -linecomment -output simpler_string.go

package consts

// Sampler 采样方法
type Sampler int

// 使用iota定义枚举值
const (
	EulerA                     Sampler = iota // 0
	Euler                      Sampler = iota // 1
	LMS                        Sampler = iota // 2
	HEUN                       Sampler = iota // 3
	DPM2                       Sampler = iota // 4
	DPM2A                      Sampler = iota // 5
	DPM2S_A_DPM_PLUS_PLUS      Sampler = iota // 6 (注意：我修改了名称以避免空格和特殊字符)
	DPM2M_DPM_PLUS_PLUS        Sampler = iota // 7
	DPM_SDE_DPM_PLUS_PLUS      Sampler = iota // 8
	DPM_FAST_DPM_PLUS_PLUS     Sampler = iota // 9
	DPM_ADAPTIVE_DPM_PLUS_PLUS Sampler = iota // 10
	LMS_Karras                 Sampler = iota // 11
	DPM2_Karras                Sampler = iota // 12
	DPM2A_Karras               Sampler = iota // 13
	// 注意：下面的DPM++ 2S a我假设你是想表达DPM++ 2S a Karras的前缀部分被省略了，
	DPM2S_A_Karras_Assumed              Sampler = iota // 14 (假设为DPM++ 2S a Karras)
	DPM2M_Karras_DPM_PLUS_PLUS          Sampler = iota // 15
	DPM_SDE_Karras_DPM_PLUS_PLUS        Sampler = iota // 16
	DDIM                                Sampler = iota // 17
	PLMS                                Sampler = iota // 18
	UNIPC                               Sampler = iota // 19
	DPM2M_SDE_Karras_DPM_PLUS_PLUS      Sampler = iota // 20
	DPM2M_SDE_EXPONENTIAL_DPM_PLUS_PLUS Sampler = iota // 21
	DPM2M_SDE_Heun_Karras_DPM_PLUS_PLUS Sampler = iota // 22 (注意：我修改了索引以匹配你的列表，但原始列表在24和25之间有跳跃)
	// 跳过23，因为原始列表中没有对应的值
	_                                        Sampler = iota
	DPM2M_SDE_Heun_Exponential_DPM_PLUS_PLUS Sampler = iota // 24 (根据原始列表的索引调整)
	DPM3M_SDE_Karras_DPM_PLUS_PLUS           Sampler = iota // 25 (我假设了3M作为下一个逻辑步骤，但请根据你的实际情况调整)
	// 跳过26，因为原始列表中没有对应的值
	_                                   Sampler = iota
	DPM3M_SDE_EXPONENTIAL_DPM_PLUS_PLUS Sampler = iota // 27
	Restart                             Sampler = iota // 28
	LCM                                 Sampler = iota // 29
)
