package packet

import (
	"strconv"
	"strings"
)

const EL = "\r\n"

// 这是一个OK等不带引号的返回
func OkLine(s string) string {
	return "+" + s + EL
}

// 带双引号的返回
func GetString(s string) string {
	return OkLine("\"" + s + "\"")
}

// 获取不带引号的大字符串
func GetBigString(s string) string {
	return "$" + s + EL
}

// 获取数组
func GetArray(ss []string) string {
	s := "*"
	return s + strings.Join(ss, EL)
}

// 错误
func ErrLine(s string) string {
	return "-" + s + EL
}

// 获取整数
func GetInteger(s int) string {
	return ":" + strconv.Itoa(s) + EL
}
