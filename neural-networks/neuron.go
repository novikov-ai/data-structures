package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

}

type Neuron struct {
	weights   [][]int
	threshold int
}

func New(n, m, threshold int) *Neuron {
	w := make([][]int, m)
	for i := range n {
		w[i] = make([]int, n)
		for j := range w[i] {
			w[i][j] = 0
		}
	}

	return &Neuron{
		weights:   w,
		threshold: threshold,
	}
}

func NewWithWeights(weights [][]int, threshold int) *Neuron {
	return &Neuron{
		weights:   weights,
		threshold: threshold,
	}
}

func (n *Neuron) Activate(file string) (bool, error) {
	data, err := os.Open(file)
	if err != nil {
		return false, err
	}

	sum, err := n.synapsisCalc(data)
	if err != nil {
		return false, err
	}

	return sum >= n.threshold, nil
}

func (n *Neuron) synapsisCalc(file *os.File) (int, error) {
	hits := 0
	i := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		nextLine := scanner.Text()
		for j, letter := range nextLine {
			v, err := strconv.Atoi(string(letter))
			if err != nil {
				return 0, err
			}

			if n.weights[i][j]*v > 0 {
				hits++
			}
		}
	}

	println(file.Name(), "count:", hits)
	return hits, nil
}
