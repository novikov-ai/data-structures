package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Activate(t *testing.T) {
	n := NewWithWeights(
		[][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		},
		35)

	tc := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "A1",
			path:     "dataset/a1.txt",
			expected: true,
		},
		{
			name:     "A2",
			path:     "dataset/a2.txt",
			expected: true,
		},
		{
			name:     "A3",
			path:     "dataset/a3.txt",
			expected: true,
		},
		{
			name:     "A4",
			path:     "dataset/a4.txt",
			expected: true,
		},
		{
			name:     "A5",
			path:     "dataset/a5.txt",
			expected: true,
		},
		{
			name:     "A6",
			path:     "dataset/a6.txt",
			expected: true,
		},
		{
			name:     "A7",
			path:     "dataset/a7.txt",
			expected: true,
		},
		{
			name:     "A8",
			path:     "dataset/a8.txt",
			expected: true,
		},
		{
			name:     "A9",
			path:     "dataset/a9.txt",
			expected: true,
		},
		{
			name:     "A10",
			path:     "dataset/a10.txt",
			expected: true,
		},
		{
			name:     "C",
			path:     "dataset/c.txt",
			expected: false,
		},
		{
			name:     "E",
			path:     "dataset/e.txt",
			expected: false,
		},
		{
			name:     "L",
			path:     "dataset/l.txt",
			expected: false,
		},
		{
			name:     "T",
			path:     "dataset/t.txt",
			expected: false,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got, err := n.Activate(tt.path)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func Test_LearnA(t *testing.T) {
	n := New(10, 10, 35)
	err := n.LearnA()
	assert.Nil(t, err)

	tc := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "A1",
			path:     "dataset/a1.txt",
			expected: true,
		},
		{
			name:     "A2",
			path:     "dataset/a2.txt",
			expected: true,
		},
		{
			name:     "A3",
			path:     "dataset/a3.txt",
			expected: true,
		},
		{
			name:     "A4",
			path:     "dataset/a4.txt",
			expected: true,
		},
		{
			name:     "A5",
			path:     "dataset/a5.txt",
			expected: true,
		},
		{
			name:     "A6",
			path:     "dataset/a6.txt",
			expected: true,
		},
		{
			name:     "A7",
			path:     "dataset/a7.txt",
			expected: true,
		},
		{
			name:     "A8",
			path:     "dataset/a8.txt",
			expected: true,
		},
		{
			name:     "A9",
			path:     "dataset/a9.txt",
			expected: true,
		},
		{
			name:     "A10",
			path:     "dataset/a10.txt",
			expected: true,
		},
		{
			name:     "C",
			path:     "dataset/c.txt",
			expected: false,
		},
		{
			name:     "E",
			path:     "dataset/e.txt",
			expected: false,
		},
		{
			name:     "L",
			path:     "dataset/l.txt",
			expected: false,
		},
		{
			name:     "T",
			path:     "dataset/t.txt",
			expected: false,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got, err := n.Activate(tt.path)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
