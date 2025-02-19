package main

type Node struct {
	value      interface{}
	parent     *Node
	left       *Node
	right      *Node
	translated string
}

func NewTree(v interface{}) *Node {
	return &Node{
		value:  v,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

func (n *Node) AddLeft(v interface{}) *Node {
	n.left = &Node{
		value:  v,
		parent: n,
		left:   nil,
		right:  nil,
	}

	return n.left
}

func (n *Node) AddRight(v interface{}) *Node {
	n.right = &Node{
		value:  v,
		parent: n,
		left:   nil,
		right:  nil,
	}

	return n.right
}

func (n *Node) AddParent(parent *Node) {
	n.parent = parent
}