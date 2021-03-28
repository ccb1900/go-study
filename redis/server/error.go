package server

import "fmt"

type Message struct {
}

func NewMessage() *Message {
	return new(Message)
}

// 参数数量错误
func (r *Message) GetWrongArgNum(s string) string {
	return fmt.Sprintf("ERR wrong number of arguments for '%s' command", s)
}

// 未授权错误
func (r *Message) GetNoAuth() string {
	return "NOAUTH Authentication required."
}

// 语法错误
func (r *Message) GetSyntaxError() string {
	return "ERR syntax error"
}

// 无效的db索引
func (r *Message) GetInvalidDbIndex() string {
	return "ERR invalid DB index"
}

// 超出范围的db索引
func (r *Message) GetOutOfRangeDbIndex() string {
	return "ERR DB index is out of range"
}
