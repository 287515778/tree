package tree

import (
	"fmt"
	"strings"
)

type node struct {
	// 路由相对路径
	pattern string
	// 当前节点的值
	part string
	// 子节点
	children []*node
}

func fetchAll(n *node) {
	for _, nn := range n.children {
		fmt.Printf("%v\n", nn)
		fetchAll(nn)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *node) insert(parts []string) {
	l := len(parts)
	for i := 0; i < l; i++ {
		child := n.matchChild(parts[i])
		if child == nil {
			child = new(node)
			child.part = parts[i]
			if i == l-1 {
				child.pattern = strings.Join(parts, "/")
			}
			fmt.Printf("child: %+v\n", child)
			n.children = append(n.children, child)
		}
		n = child
	}
}

func (n *node) search(parts []string) *node {
	for _, part := range parts {
		child := n.matchChild(part)
		if child == nil {
			return nil
		}
		n = child
	}
	if n.pattern != "" {
		return n
	}
	return nil
}
