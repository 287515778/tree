package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	rootNode := &node{}
	parts := []string{"go", "api", "get"}
	rootNode.insert(parts)

	parts1 := []string{"go", "api", "post"}
	rootNode.insert(parts1)

	fetchAll(rootNode)

	n := rootNode.search(parts1)
	fmt.Printf("search node: %v\n", n)
}
