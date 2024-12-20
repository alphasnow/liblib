//go:generate stringer -type=GenerateStatus -linecomment -output generate_status_string.go

package consts

type GenerateStatus int

const (
	Waiting   GenerateStatus = iota + 1 // 1:等待执行
	Executing                           // 2:执行中
	Generated                           // 3:已生图
	Reviewing                           // 4:审核中
	Success                             // 5:成功
	Failure                             // 6:失败
)

const (
	Expired GenerateStatus = iota + 10 // 10:任务已过期(自定义)
)
