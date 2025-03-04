package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	ID                    int
}

type Leaf struct {
	Tree1 *Node
	Tree2 *Node
	Color int
}

type BWT struct {
	Tree1     *Node
	Tree2     *Node
	LeafLinks []Leaf
	Depth     int
	NumColors int
	Symmetric bool
}

type Path struct {
	Colors []int  
	Leaf   *Node
	Count  int    
}

func NewBWT(depth, numColors, startID int, symmetric bool) *BWT {
	rand.Seed(time.Now().UnixNano())

	nextID1 := startID
	tree1 := buildTree(depth, numColors, &nextID1)

	var tree2 *Node
	nextID2 := nextID1
	if symmetric {
		tree2 = buildMirrorTree(tree1, &nextID2)
	} else {
		tree2 = buildTree(depth, numColors, &nextID2)
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

func buildTree(depth, numColors int, nextID *int) *Node {
	if depth <= 0 {
		return nil
	}

	node := &Node{ID: *nextID}
	(*nextID)++

	assignColors(node, numColors)
	node.Left = buildTree(depth-1, numColors, nextID)
	node.Right = buildTree(depth-1, numColors, nextID)
	return node
}

func buildMirrorTree(src *Node, nextID *int) *Node {
	if src == nil {
		return nil
	}

	node := &Node{
		ColorLeft:  src.ColorRight,
		ColorRight: src.ColorLeft,
		ID:         *nextID,
	}
	(*nextID)++

	node.Left = buildMirrorTree(src.Left, nextID)
	node.Right = buildMirrorTree(src.Right, nextID)
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

	var leaves []*Node
	leaves = append(leaves, getLeaves(node.Left)...)
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

// Сложность алгоритма экспоненциальная!
// (O(D*2^D))
func (b *BWT) FindOptimalPaths(targetColor int, minimize bool) ([]Path, []Path) {
	var paths1, paths2 []Path
	collectPaths(b.Tree1, []int{}, &paths1)
	collectPaths(b.Tree2, []int{}, &paths2)

	for i := range paths1 {
		paths1[i].Count = countColor(paths1[i].Colors, targetColor)
	}
	for i := range paths2 {
		paths2[i].Count = countColor(paths2[i].Colors, targetColor)
	}

	sortPaths(paths1, minimize)
	sortPaths(paths2, minimize)

	return paths1, paths2
}

func collectPaths(node *Node, currentColors []int, paths *[]Path) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, Path{
			Colors: append([]int{}, currentColors...),
			Leaf:   node,
		})
		return
	}

	if node.Left != nil {
		collectPaths(node.Left, append(currentColors, node.ColorLeft), paths)
	}
	if node.Right != nil {
		collectPaths(node.Right, append(currentColors, node.ColorRight), paths)
	}
}

func countColor(colors []int, target int) int {
	count := 0
	for _, c := range colors {
		if c == target {
			count++
		}
	}
	return count
}

func sortPaths(paths []Path, minimize bool) {
	sort.Slice(paths, func(i, j int) bool {
		if minimize {
			return paths[i].Count < paths[j].Count
		}
		return paths[i].Count > paths[j].Count
	})
}

func printTreeColored(node *Node, prefix string, isTail bool, branchColor string) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	fmt.Printf("{%v}", node.ID)
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
	fmt.Printf("{%v}", node.ID)
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

// Случайное блуждание требует примерно O(√N) шагов для достижения цели.
func (b *BWT) RandomWalk(targetID int) (*Node, int) {
	current := b.Tree1
	steps := 0

	for current.ID != targetID {
		steps++
		move := rand.Intn(3) // 0 - влево, 1 - вправо, 2 - случайный переход между деревьями

		switch move {
		case 0:
			if current.Left != nil {
				current = current.Left
			}
		case 1:
			if current.Right != nil {
				current = current.Right
			}
		case 2:
			for _, link := range b.LeafLinks {
				if link.Tree1 == current {
					current = link.Tree2
					break
				} else if link.Tree2 == current {
					current = link.Tree1
					break
				}
			}
		}
	}

	return current, steps
}

// Квантовое блуждание находит цель быстрее, но требует управления экспоненциальным ростом состояний.
func (b *BWT) QuantumWalk(targetID int, maxStates int) (*Node, int) {
	type State struct {
		Node  *Node
		Steps int
	}

	states := []State{{Node: b.Tree1, Steps: 0}}
	visited := make(map[int]bool)

	for len(states) > 0 {
		newStates := []State{}
		for _, state := range states {
			if state.Node.ID == targetID {
				return state.Node, state.Steps
			}

			if !visited[state.Node.ID] {
				visited[state.Node.ID] = true

				if state.Node.Left != nil {
					newStates = append(newStates, State{Node: state.Node.Left, Steps: state.Steps + 1})
				}
				if state.Node.Right != nil {
					newStates = append(newStates, State{Node: state.Node.Right, Steps: state.Steps + 1})
				}
				for _, link := range b.LeafLinks {
					if link.Tree1 == state.Node {
						newStates = append(newStates, State{Node: link.Tree2, Steps: state.Steps + 1})
					} else if link.Tree2 == state.Node {
						newStates = append(newStates, State{Node: link.Tree1, Steps: state.Steps + 1})
					}
				}
			}
		}

		// Ограничение количества состояний
		if len(newStates) > maxStates {
			rand.Shuffle(len(newStates), func(i, j int) { newStates[i], newStates[j] = newStates[j], newStates[i] })
			newStates = newStates[:maxStates]
		}

		states = newStates
	}

	return nil, -1 // Не найден
}
