package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTokens(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []Node
	}{
		{
			name:  "simple",
			input: "((7+3)*(5-2))",
			output: []Node{
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "число",
					TokenValue: "7",
				},
				{
					TokenType:  "операция",
					TokenValue: "+",
				},
				{
					TokenType:  "число",
					TokenValue: "3",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
				},
				{
					TokenType:  "операция",
					TokenValue: "*",
				},
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "число",
					TokenValue: "5",
				},
				{
					TokenType:  "операция",
					TokenValue: "-",
				},
				{
					TokenType:  "число",
					TokenValue: "2",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
				},
			},
		},
		{
			name:  "with brackets wrapping",
			input: "7+3/25*(5-2)",
			output: []Node{
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "число",
					TokenValue: "7",
				},
				{
					TokenType:  "операция",
					TokenValue: "+",
				},
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "число",
					TokenValue: "3",
				},
				{
					TokenType:  "операция",
					TokenValue: "/",
				},
				{
					TokenType:  "число",
					TokenValue: "2",
				},
				{
					TokenType:  "число",
					TokenValue: "5",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
				},
				{
					TokenType:  "операция",
					TokenValue: "*",
				},
				{
					TokenType:  "скобка",
					TokenValue: "(",
				},
				{
					TokenType:  "число",
					TokenValue: "5",
				},
				{
					TokenType:  "операция",
					TokenValue: "-",
				},
				{
					TokenType:  "число",
					TokenValue: "2",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
				},
				{
					TokenType:  "скобка",
					TokenValue: ")",
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
