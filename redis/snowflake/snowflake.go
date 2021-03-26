package snowflake

var count int

func GetId() int {
	count++
	return count
}
