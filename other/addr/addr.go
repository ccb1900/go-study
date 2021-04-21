package addr

import (
	"fmt"
	"strconv"
	"strings"
)

func test(a *int) {
	fmt.Printf("%v\n", &a)
}

func insert(root *Node, node *Node) {
	if len(node.Type)-len(root.Type) == 1 {
		root.Next = append(root.Next, node)
	} else {
		//insert()
	}
}

type Node struct {
	Text   string
	Type   string
	Prefix string
	Index  int
	Next   []*Node
}

func initNode(token string) *Node {
	tokens := strings.Split(token, " ")
	return &Node{
		Text: tokens[1],
		Type: tokens[0],
		Next: make([]*Node, 0),
	}
}

type Tree struct {
	Root *Node
}

func translate(data []string) [][]string {
	status := ""
	prefix := ""
	index := 0
	for _, datum := range data {
		tokens := strings.Split(datum, " ")
		if len(tokens[0]) > len(status) {
			index++
			status = tokens[0]
		}
		if tokens[0] == status {
			prefix += strconv.Itoa(index)
			index = 0
			index++
		}

		if len(tokens[0]) < len(status) {

		}
	}

	return nil
}
