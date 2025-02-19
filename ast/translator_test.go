package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Translate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "ok",
			input:  "(7+((3*5)-2))",
			output: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ast := NewAST(tt.input)
			node := ast.Create()

			got, ok := getIntValue(Translate(node))
			assert.True(t, ok)
			assert.Equal(t, tt.output, got)
		})
	}
}
