package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ansiColors = []string{
	"\033[31m", // красный
	"\033[32m", // зеленый
	"\033[33m", // желтый
	"\033[34m", // синий
	"\033[35m", // магента
	"\033[36m", // циан
}

func colorize(c int, text string) string {
	color := ansiColors[c%len(ansiColors)]
	return color + text + "\033[0m"
}

type Node struct {
	Left, Right           *Node
	ColorLeft, ColorRight int
}

// Leaf связывает листья двух деревьев и хранит цвет соединяющей ветви.
type Leaf struct {
	Tree1 *Node
	Tree2 *Node
	Color int
}

// BWT представляет двоичное спаянное (склеенное) дерево.
type BWT struct {
	Tree1     *Node
	Tree2     *Node
	LeafLinks []Leaf
	Depth     int
	NumColors int
	Symmetric bool
}

// Если symmetric == true, второе дерево создаётся зеркальным по цветам относительно первого.
func NewBWT(depth, numColors int, symmetric bool) *BWT {
	rand.Seed(time.Now().UnixNano())

	tree1 := buildTree(depth, numColors)
	var tree2 *Node
	if symmetric {
		tree2 = buildMirrorTree(tree1)
	} else {
		tree2 = buildTree(depth, numColors)
	}

	bwt := &BWT{
		Tree1:     tree1,
		Tree2:     tree2,
		Depth:     depth,
		NumColors: numColors,
		Symmetric: symmetric,
	}
	bwt.linkLeaves()
	return bwt
}

func buildTree(depth, numColors int) *Node {
	if depth <= 0 {
		return nil
	}
	node := &Node{}
	assignColors(node, numColors)
	node.Left = buildTree(depth-1, numColors)
	node.Right = buildTree(depth-1, numColors)
	return node
}

func buildMirrorTree(src *Node) *Node {
	if src == nil {
		return nil
	}
	node := &Node{
		ColorLeft:  src.ColorRight,
		ColorRight: src.ColorLeft,
	}
	node.Left = buildMirrorTree(src.Left)
	node.Right = buildMirrorTree(src.Right)
	return node
}

func assignColors(node *Node, numColors int) {
	colorLeft := rand.Intn(numColors)
	colorRight := rand.Intn(numColors)
	for colorRight == colorLeft {
		colorRight = rand.Intn(numColors)
	}
	node.ColorLeft = colorLeft
	node.ColorRight = colorRight
}

func getLeaves(node *Node) []*Node {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []*Node{node}
	}
	leaves := getLeaves(node.Left)
	leaves = append(leaves, getLeaves(node.Right)...)
	return leaves
}

func (b *BWT) linkLeaves() {
	leaves1 := getLeaves(b.Tree1)
	leaves2 := getLeaves(b.Tree2)

	rand.Shuffle(len(leaves2), func(i, j int) { leaves2[i], leaves2[j] = leaves2[j], leaves2[i] })

	for i := 0; i < len(leaves1) && i < len(leaves2); i++ {
		color := rand.Intn(b.NumColors)
		b.LeafLinks = append(b.LeafLinks, Leaf{
			Tree1: leaves1[i],
			Tree2: leaves2[i],
			Color: color,
		})
	}
}

func printTreeColored(node *Node, prefix string, isTail bool, branchColor string) {
	if node == nil {
		return
	}
	fmt.Print(prefix)
	if branchColor != "" {
		fmt.Print(branchColor, "→")
	}
	if isTail {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}
	fmt.Println("●")
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		rightBranchColor := colorize(node.ColorRight, fmt.Sprintf("%d", node.ColorRight))
		printTreeColored(node.Right, newPrefix, false, rightBranchColor)
	}
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		leftBranchColor := colorize(node.ColorLeft, fmt.Sprintf("%d", node.ColorLeft))
		printTreeColored(node.Left, newPrefix, true, leftBranchColor)
	}
}

func printTreeColoredInverted(node *Node, prefix string, isTail bool, branchColor string) {
	if node == nil {
		return
	}
	fmt.Print(prefix)
	if branchColor != "" {
		fmt.Print(branchColor, "→")
	}
	if isTail {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}
	fmt.Println("●")
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		leftBranchColor := colorize(node.ColorLeft, fmt.Sprintf("%d", node.ColorLeft))
		printTreeColoredInverted(node.Left, newPrefix, false, leftBranchColor)
	}
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		rightBranchColor := colorize(node.ColorRight, fmt.Sprintf("%d", node.ColorRight))
		printTreeColoredInverted(node.Right, newPrefix, true, rightBranchColor)
	}
}

func (b *BWT) Print() {
	fmt.Println("=== Tree 1 ===")
	printTreeColored(b.Tree1, "", true, "")

	fmt.Println("\n=== Leaf Connections ===")
	for i, link := range b.LeafLinks {
		connColor := colorize(link.Color, fmt.Sprintf("%d", link.Color))
		fmt.Printf("Соединение %d: Лист Tree1 (%p) —%s— Лист Tree2 (%p)\n",
			i+1, link.Tree1, connColor, link.Tree2)
	}

	fmt.Println("\n=== Tree 2 (Inverted) ===")
	printTreeColoredInverted(b.Tree2, "", true, "")
}