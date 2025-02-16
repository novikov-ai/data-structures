package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ast(t *testing.T) {
	tests := []struct {
		name     string
		expected *Node
	}{
		{
			name: "ok",
			expected: createTree(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ast := NewAST("(7+((3*5)-2))")
			got := ast.Create()
			assert.Equal(t, tt.expected, got)

			PrintTree(got)
			PrintTree(tt.expected)
		})
	}
}

func createTree() *Node {
	nodePlus := NewTree("+")

	nodePlus.AddLeft(7)

	nodeMinus := nodePlus.AddRight("-")

	asteriskNode := nodeMinus.AddLeft("*")
	nodeMinus.AddRight(2)

	asteriskNode.AddLeft(3)
	asteriskNode.AddRight(5)

	return nodePlus
}

func PrintTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	printSubtree(root, "", true)
}

func printSubtree(node *Node, prefix string, isTail bool) {
	if node == nil {
		return
	}

	children := []*Node{}
	if node.left != nil {
		children = append(children, node.left)
	}
	if node.right != nil {
		children = append(children, node.right)
	}

	for i, child := range children {
		isLast := i == len(children)-1
		
		connector := "├── "
		nextPrefix := prefix + "│   "
		if isLast {
			connector = "└── "
			nextPrefix = prefix + "    "
		}
		
		// Print current child
		fmt.Printf("%s%s%v\n", prefix, connector, child.value)
		
		// Recursively print child's subtree
		printSubtree(child, nextPrefix, isLast)
	}
}