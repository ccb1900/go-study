package node

type EL int

// Single 单节点定义
type Single struct {
	Data EL
	Next *Single
}

// CreateSingle 创建新的单节点
func CreateSingle(data EL) *Single {
	n := new(Single)
	n.Data = data
	return n
}
