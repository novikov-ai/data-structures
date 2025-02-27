package main

func main() {
	// Пример: глубина 3, 4 цвета, включён режим зеркальной раскраски (symmetric = true)
	bwt := NewBWT(3, 4, true)
	bwt.Print()
}
