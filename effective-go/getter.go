package effective_go

type Who struct {
	owner string
}

// 不推荐使用GetOwner
// 获取器
func (w *Who) Owner() string {
	return w.owner
}

// 设置器
func (w *Who) SetOwner(owner string) {
	w.owner = owner
}
func (w *Who) testError(a int) (int, error) {
	oE := &OwnerError{
		name:  w.owner,
		value: w.owner,
		Err:   nil,
	}

	if a > 10 {
		return 0, nil
	}
	return a, oE
}

// 标准的字符串转换方法名称，实现了对应的内置接口，（w Who）调用和声明必须一致，声明指针就必须由指针调用，原生的就必须是原生的调用
func (w *Who) String() string {
	return "字符串转换方法" + w.owner
}
