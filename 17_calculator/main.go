package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	if len(tokens) == 0 {
		return 0, errors.New("empty expression")
	}

	rpn, err := toRPN(tokens)
	if err != nil {
		return 0, err
	}

	return evaluateRPN(rpn)
}

func tokenize(expression string) []string {
	expression = strings.ReplaceAll(expression, " ", "")
	tokens := []string{}
	number := ""

	for _, char := range expression {
		switch {
		case char >= '0' && char <= '9' || char == '.':
			number += string(char)
		case char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')':
			if number != "" {
				tokens = append(tokens, number)
				number = ""
			}
			tokens = append(tokens, string(char))
		default:
			return nil
		}
	}
	if number != "" {
		tokens = append(tokens, number)
	}
	return tokens
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func toRPN(tokens []string) ([]string, error) {
	output := []string{}
	operators := []string{}

	for _, token := range tokens {
		if isOperator(token) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
		} else {
			output = append(output, token)
		}
	}
	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return output, nil
}

func evaluateRPN(tokens []string) (float64, error) {
	stack := []float64{}

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			}
			stack = append(stack, result)
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("invalid number")
			}
			stack = append(stack, num)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}

func main() {
	expressions := []string{
		"5 + 7 * 3 / (2 - 6) * 4 + 1",
		"15 + (4 * 6)",
		"200 / 4 * 3 + 10",
		"8 + 6 * 4 / (2 - 6) ^ 3 ^ 2",
	}

	for _, expr := range expressions {
		result, err := Calc(expr)
		if err != nil {
			fmt.Printf("Error in expression \"%s\": %v\n", expr, err)
		} else {
			fmt.Printf("Result of \"%s\" = %f\n", expr, result)
		}
	}
}
