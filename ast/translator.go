package main

import (
	"fmt"
	"strconv"
)

func Translate(n *Node) interface{} {
	v, ok := n.value.(string)
	if !ok && n.parent == nil {
		return n.value
	}

	var leftValue, rightValue int

	if n.left != nil {
		lf, ok := getIntValue(n.left.value)
		if !ok {
			leftValue, _ = getIntValue(Translate(n.left))
		} else {
			leftValue = lf
		}
	}

	if n.right != nil {
		rf, ok := getIntValue(n.right.value)
		if !ok {
			rightValue, _ = getIntValue(Translate(n.right))
		} else {
			rightValue = rf
		}
	}

	switch v {
	case "-":
		n.value = leftValue - rightValue
	case "+":
		n.value = leftValue + rightValue
	case "/":
		n.value = leftValue / rightValue
	case "*":
		n.value = leftValue * rightValue
	}

	n.translated = fmt.Sprintf("(%v%s%v)", leftValue, v, rightValue)

	n.left = nil
	n.right = nil

	return n.value
}

func getIntValue(v interface{}) (int, bool) {
	intValue, ok := v.(int)
	if ok {
		return intValue, true
	}

	value, ok := v.(string)
	if !ok {
		return 0, false
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}

	return intValue, true
}