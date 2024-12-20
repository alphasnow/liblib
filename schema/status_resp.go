package schema

type StatusResp struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data StatusData `json:"data"`
}

type StatusData struct {
	GenerateUuid     string        `json:"generateUuid" example:"8dcbfa2997444899b71357ccb7db378b"` // 生图任务uuid，使用该uuid查询生图进度
	GenerateStatus   int           `json:"generateStatus" example:"5"`                              // 1等待执行,2执行中,3已生图,4审核中,5成功,6失败
	PercentCompleted float64       `json:"percentCompleted" example:"1"`                            // 生图进度，0到1之间的浮点数，暂未实现生图进度
	GenerateMsg      string        `json:"generateMsg" example:""`                                  // 生图信息，提供附加信息，如生图失败信息
	Images           []StatusImage `json:"images"`                                                  // 图片列表，只提供审核通过的图片
}

type StatusImage struct {
	ImageUrl    string `json:"imageUrl" example:"https://liblibai-tmp-image.liblib.cloud/img/b10c2631c3024e2aae0816a17867c679/1c82d6fc708d51e16ce10f5afcc1763af88a3d803381555c7725a152a2f15e4c.png"` // 图片地址，可直接访问，地址有时效性：7天
	Seed        int64  `json:"seed" example:"299499720"`                                                                                                                                             // 随机种子值
	AuditStatus int    `json:"auditStatus" example:"3"`                                                                                                                                              // 1待审核,2审核中,3审核通过,4审核拦截,5审核失败
}
