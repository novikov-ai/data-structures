package main

func main() {
	// Пример: глубина 3, 4 цвета, включён режим зеркальной раскраски (symmetric = true)
	bwt := NewBWT(3, 4, 0, true)
	bwt.Print()

	s1,s2 := bwt.FindOptimalPaths(1, false)
	print("%v%v", s1,s2)
}
