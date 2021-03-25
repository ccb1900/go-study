package effective_go

// 自定义错误
type OwnerError struct {
	name  string // 自定义信息
	value string // 自定义信息
	Err   error  // 系统错误
}

// 实现了内置错误接口
func (oE *OwnerError) Error() string {
	s := "i do not know what happen " + oE.name + "-value " + oE.value
	if oE.Err == nil {
		return s + "没有提供error信息"
	}
	return s + oE.Err.Error()
}
