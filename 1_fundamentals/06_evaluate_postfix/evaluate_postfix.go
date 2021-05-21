package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 1.3.11 Write a program EvaluatePostfix that takes a postfix expression from standard input, evaluates it, and prints the value.
// (Piping the output of your program from the previous exercise to this program gives equivalent behavior to Evaluate.)
// 234+56**+
// 212
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(line[:len(line)-1], "")

	stack := &Stack{}
	for _, s := range str {
		if s == "+" {
			stack.push(stack.pop()+stack.pop())
		} else if s == "*" {
			stack.push(stack.pop()*stack.pop())
		} else {
			i,_ := strconv.Atoi(s)
			stack.push(i)
		}
	}
	fmt.Println(stack.pop())
}

type Stack []int

func (stack *Stack) push(arr int) {
	*stack = append(*stack, arr)
}

func (stack *Stack) pop() int {
	s := *stack
	last := s[len(s)-1]
	*stack = s[:len(s)-1]
	return last
}

func (stack *Stack) peek() int {
	s := *stack
	return s[len(s)-1]
}

func (stack *Stack) size() int {
	return len(*stack)
}
