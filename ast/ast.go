package main

type AST struct {
	tokenList []Token
	node      *Node
}

func NewAST(v string) AST {
	return AST{
		tokenList: GetTokens(v),
		node:      &Node{},
	}
}

func (a *AST) Create() *Node {
	if len(a.tokenList) == 0 {
		return nil
	}

	root := a.node

	for _, n := range a.tokenList {
		switch n.Type {
		case "скобка":
			if n.Value == "(" {
				empty := &Node{
					parent: a.node,
				}

				a.node = a.node.AddLeft(empty)
			}
			if n.Value == ")" {
				a.node = a.node.parent
			}
		case "число":
			a.node.value = n.Value
			a.node = a.node.parent
		case "операция":
			a.node.value = n.Value
			empty := &Node{
				parent: a.node,
			}

			a.node = a.node.AddRight(empty)
		}
	}

	return root
}