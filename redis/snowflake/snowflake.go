package snowflake

var count int

// 生成客户端id
func GetId() int {
	count++
	return count
}
