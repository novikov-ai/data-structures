package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

}

var dataSet = map[string]rune{
	"dataset/a1.txt":  'a',
	"dataset/a2.txt":  'a',
	"dataset/a3.txt":  'a',
	"dataset/a4.txt":  'a',
	"dataset/a5.txt":  'a',
	"dataset/a6.txt":  'a',
	"dataset/a7.txt":  'a',
	"dataset/a8.txt":  'a',
	"dataset/a9.txt":  'a',
	"dataset/a10.txt": 'a',
	"dataset/c.txt":   'c',
	"dataset/e.txt":   'e',
	"dataset/l.txt":   'l',
	"dataset/t.txt":   't',
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

func (n *Neuron) LearnA() error {
	for path, v := range dataSet {
		ok, err := n.Activate(path)
		if err != nil {
			return err
		}

		if rune(v) == 'a' {
			if ok {
				continue
			}

			n.IncreaseWeights(path)
			continue
		}

		if ok {
			n.DecreaseWeights(path)
		}
	}

	return nil
}

func (n *Neuron) IncreaseWeights(file string) error {
	data, err := os.Open(file)
	if err != nil {
		return err
	}

	return n.changeWeight(data, func(value int) int {
		return value * (1)
	})
}

func (n *Neuron) DecreaseWeights(file string) error {
	data, err := os.Open(file)
	if err != nil {
		return err
	}

	return n.changeWeight(data, func(value int) int {
		return value * (-1)
	})
}

func (n *Neuron) changeWeight(file *os.File, action func(value int) int) error {
	i := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		nextLine := scanner.Text()
		for j, letter := range nextLine {
			v, err := strconv.Atoi(string(letter))
			if err != nil {
				return err
			}

			n.weights[i][j] += action(v)
		}
	}

	return nil
}
