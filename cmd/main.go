package main

import (
	"bufio"
	"fmt"
	"os"

	scary "github.com/jackdoe/go-scary/pkg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(scary.Scary(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
