package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTokens(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []Token
	}{
		{
			name:  "simple",
			input: "((7+3)*(5-2))",
			output: []Token{
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "7",
				},
				{
					Type:  "операция",
					Value: "+",
				},
				{
					Type:  "число",
					Value: "3",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "операция",
					Value: "*",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "5",
				},
				{
					Type:  "операция",
					Value: "-",
				},
				{
					Type:  "число",
					Value: "2",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
			},
		},
		{
			name:  "with brackets wrapping",
			input: "7+3/25*(5-2)",
			output: []Token{
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "7",
				},
				{
					Type:  "операция",
					Value: "+",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "3",
				},
				{
					Type:  "операция",
					Value: "/",
				},
				{
					Type:  "число",
					Value: "2",
				},
				{
					Type:  "число",
					Value: "5",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "операция",
					Value: "*",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "5",
				},
				{
					Type:  "операция",
					Value: "-",
				},
				{
					Type:  "число",
					Value: "2",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
			},
		},
		{
			name: "with brackets wrapping v2",
			input: "7+3*5-2", 
			output: []Token{
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "7",
				},
				{
					Type:  "операция",
					Value: "+",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "скобка",
					Value: "(",
				},
				{
					Type:  "число",
					Value: "3",
				},
				{
					Type:  "операция",
					Value: "*",
				},
				{
					Type:  "число",
					Value: "5",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "операция",
					Value: "-",
				},
				{
					Type:  "число",
					Value: "2",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
				{
					Type:  "скобка",
					Value: ")",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, GetTokens(tt.input))
		})
	}
}
