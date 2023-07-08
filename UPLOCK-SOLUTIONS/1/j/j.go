package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    // Read input lines until there are no more
    for input.Scan() {
        counts[input.Text()]++
    }

    // Print lines that appear more than once
    for line, n := range counts {
        if n > 0 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

