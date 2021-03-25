package effective_go

// new不会初始化内存，只会设定零值(递归设定)，返回对应类型的指针，所以和var不同
// 变量与内存的关系，var声明，只是说需要多少空间，没有实际分配。赋值是分配内存并设置值。new 是分配内存，并设置零值（递归设定，所以属性都会被设置），返回地址
// 零值：布尔的零值是false，数字是0，字符串是""
// make 返回类型为T的已初始化的值（非零值），支持且支持切片，channel，map，这三种类型为引用类型，使用前必须初始化，不能用于其他基本类型
// 切片是一个具有三项内容的描述符，包含一个指向数据的指针，长度以及容量，在被初始化之前，切片为nil
// 同样可以new([]int)切片，但是结果是零值，也就是nil
// nil for pointers, functions, interfaces, slices, channels, and maps
// 所以new和var还是不一样的
// https://golang.org/ref/spec#The_zero_value
// new 可以为结构体的属性设置零值
// map 无法使用new，其他两种使用new都得到nil，接口也是nil，结构体会返回地址并为属性设定零值，其他基本类型都是返回地址，设定零值

type EmptyData struct {
}
type Data struct {
	Name string
	Val  IData
	F    func() string
}
type IData interface {
}
