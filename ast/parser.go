package main

import (
	"strconv"
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

func GetTokens(vv string) []Token {
	vv = wrapWithBrackets(vv)

	parsed := make([]Token, 0, len(vv))
	for _, v := range vv {
		for _, rule := range rules {
			value := string(v)
			token, ok := rule(value)
			if !ok {
				continue
			}
			parsed = append(parsed, Token{
				Type:  token,
				Value: value,
			})
		}
	}

	return parsed
}

func wrapWithBrackets(vv string) string {
	p := newParser(vv)
	expr := p.parseExpression()
	return toString(expr)
}

var rules = []func(v string) (string, bool){
	func(v string) (string, bool) {
		return "скобка", v == ")" || v == "("
	},
	func(v string) (string, bool) {
		_, err := strconv.Atoi(v)
		return "число", err == nil
	},
	func(v string) (string, bool) {
		return "операция", v == "/" || v == "*" || v == "-" || v == "+"
	},
}

type Expr interface{}

type Number struct {
	value string
}

type Binary struct {
	op    rune
	left  Expr
	right Expr
}

type parser struct {
	input string
	pos   int
}

func newParser(input string) *parser {
	return &parser{input: input, pos: 0}
}

func (p *parser) peek() rune {
	if p.pos < len(p.input) {
		return rune(p.input[p.pos])
	}
	return 0
}

func (p *parser) next() rune {
	ch := p.peek()
	p.pos++
	return ch
}

func (p *parser) eatWhitespace() {
	for p.pos < len(p.input) && unicode.IsSpace(rune(p.input[p.pos])) {
		p.pos++
	}
}

func (p *parser) parseExpression() Expr {
	node := p.parseTerm()
	p.eatWhitespace()
	for p.pos < len(p.input) {
		p.eatWhitespace()
		ch := p.peek()
		if ch == '+' || ch == '-' {
			op := p.next()
			p.eatWhitespace()
			right := p.parseTerm()
			node = &Binary{op: op, left: node, right: right}
		} else {
			break
		}
	}
	return node
}

func (p *parser) parseTerm() Expr {
	node := p.parseFactor()
	p.eatWhitespace()
	for p.pos < len(p.input) {
		p.eatWhitespace()
		ch := p.peek()
		if ch == '*' || ch == '/' {
			op := p.next()
			p.eatWhitespace()
			right := p.parseFactor()
			node = &Binary{op: op, left: node, right: right}
		} else {
			break
		}
	}
	return node
}

func (p *parser) parseFactor() Expr {
	p.eatWhitespace()
	if p.pos < len(p.input) && p.peek() == '(' {
		p.next()
		p.eatWhitespace()
		node := p.parseExpression()
		p.eatWhitespace()
		if p.pos < len(p.input) && p.peek() == ')' {
			p.next()
		} else {
			panic("expected closing parenthesis")
		}
		return node
	}
	start := p.pos
	for p.pos < len(p.input) && (unicode.IsDigit(rune(p.input[p.pos])) || p.input[p.pos] == '.') {
		p.pos++
	}
	return &Number{value: p.input[start:p.pos]}
}

func toString(expr Expr) string {
	switch e := expr.(type) {
	case *Number:
		return e.value
	case *Binary:
		return "(" + toString(e.left) + string(e.op) + toString(e.right) + ")"
	}
	return ""
}
