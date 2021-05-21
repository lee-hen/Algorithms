package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 1.3.4 Write a stack client Parentheses that reads in a text stream from standard input
// and uses a stack to determine whether its parentheses are properly balanced.
// For example, your program should print true for [()]{}{[()()]()} and false for [(]).

const (
	leftParen = '('
	rightParen = ')'
	leftBrace = '{'
	rightBrace ='}'
	leftBracket = '['
	rightBracket = ']'
)

func IsBalanced(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == leftParen {
			stack = append(stack, leftParen)
		}

		if s[i] == leftBrace {
			stack = append(stack, leftBrace)
		}

		if s[i] == leftBracket {
			stack = append(stack, leftBracket)
		}
		
		if s[i] == rightParen {
			if len(stack) == 0 {
				return false
			}

			peek := stack[len(stack)-1]
			if peek != leftParen {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if s[i] == rightBrace {
			if len(stack) == 0 {
				return false
			}

			peek := stack[len(stack)-1]
			if peek != leftBrace {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if s[i] == rightBracket {
			if len(stack) == 0 {
				return false
			}

			peek := stack[len(stack)-1]
			if peek != leftBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		
	}

	return len(stack) == 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(IsBalanced(str))
}
