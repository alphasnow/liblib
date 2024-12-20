package schema

type GenerateResp struct {
	Code int          `json:"code"` // 成功:0
	Msg  string       `json:"msg"`  // 消息
	Data GenerateData `json:"data"` // 负载
}

type GenerateData struct {
	GenerateUuid string `json:"generateUuid"`
}
