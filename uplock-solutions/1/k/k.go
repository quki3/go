package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countLines(os.Stdin, counts)
	}else{
		for _, arg := range file {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fmt.Println("output os.Args[1:]")
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
		fmt.Println(counts)
	}
}
