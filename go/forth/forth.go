package forth

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func handleWord(words map[string]string, name string, instructions string) error {
	if _, err := strconv.Atoi(name); err == nil {
		return errors.New("illegal operation")
	}
	for word, instruction := range words {
		r := regexp.MustCompile(`(?i)` + word + `+`)
		if r.MatchString(instructions) {
			instructions = r.ReplaceAllString(instructions, instruction)
		}
	}
	words[name] = instructions
	return nil
}

func Forth(input []string) ([]int, error) {
	var stack []int
	words := make(map[string]string)
	for _, line := range input {
		reWord := regexp.MustCompile(`: ([a-zA-Z+\-*/]+) (.+) ;`)
		if parse := reWord.FindAllStringSubmatch(line, 1); parse != nil {
			err := handleWord(words, strings.ToLower(parse[0][1]), parse[0][2])
			if err != nil {
				return stack, err
			}
		} else {
			parsedInput := strings.Split(line, " ")
			i := 0
			for i < len(parsedInput) {
				item := parsedInput[i]
				reNum := regexp.MustCompile(`[0-9]+`)
				reOp := regexp.MustCompile(`^[+\-*/]$`)
				reManipulation := regexp.MustCompile(`(?i)^(dup|drop|swap|over)$`)
				if reNum.MatchString(item) {
					value, err := strconv.Atoi(item)
					if err != nil {
						return stack, err
					}
					stack = append(stack, value)
				} else if instructions, ok := words[strings.ToLower(item)]; ok {
					parsedInput = append(parsedInput, strings.Split(instructions, " ")...)
				} else if reOp.MatchString(item) {
					if len(stack) == 0 {
						return stack, errors.New("empty stack")
					}
					if len(stack) == 1 {
						return stack, errors.New("only one value on the stack")
					}
					n1 := stack[len(stack)-2]
					n2 := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					switch item {
					case "+":
						stack[len(stack)-1] = n1 + n2
					case "-":
						stack[len(stack)-1] = n1 - n2
					case "*":
						stack[len(stack)-1] = n1 * n2
					case "/":
						if n2 == 0 {
							return stack, errors.New("divide by zero")
						}
						stack[len(stack)-1] = n1 / n2
					}
				} else if reManipulation.MatchString(item) {
					switch strings.ToLower(item) {
					case "swap":
						if len(stack) < 2 {
							return stack, errors.New("missing number")
						}
						temp := stack[len(stack)-1]
						stack[len(stack)-1] = stack[len(stack)-2]
						stack[len(stack)-2] = temp
					case "dup":
						if len(stack) < 1 {
							return stack, errors.New("missing number")
						}
						stack = append(stack, stack[len(stack)-1])
					case "over":
						if len(stack) < 2 {
							return stack, errors.New("missing number")
						}
						stack = append(stack, stack[len(stack)-2])
					case "drop":
						if len(stack) < 1 {
							return stack, errors.New("missing number")
						}
						stack = stack[:len(stack)-1]
					}
				} else {
					return stack, errors.New("undefined operation")
				}
				i++
			}
		}
	}
	return stack, nil
}
