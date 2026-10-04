package random

import (
	"bufio"
	"fmt"
	"io"
)

func readContent(r io.Reader) {
	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("Somsthing went wrong")
		return
	}
	fmt.Println(string(data))
}

func readLines(r io.Reader) {
	scanner := bufio.NewScanner(r)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading failed:", err)
	}
}
