package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 1.3.10 Write a filter InfixToPostfix that converts an arithmetic expression from infix to postfix.
//(2+((3+4)*(5*6)))
//2 3 4 + 5 6 * * +
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

	stack := make([]string, 0)
	for _, s := range str {
		if s == "+" {
			stack = append(stack, s)
		} else if s == "*" {
			stack = append(stack, s)
		} else if s == ")" {
			peek := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Print(peek + " ")
		}  else if s == "(" {
			fmt.Print("")
		} else {
			fmt.Print(s + " ")
		}
	}
}

