package main
// bufio - provides a convenient set of functions and types for reading and writing data efficiently.
// fmt - provides basic I/O functionality, including formatted input and output.
// os - provides a platform-independent interface to operating system functionality. 
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
            // make - function is a built-in function used to create and initialize certain types of data structures, such as slices, maps, and channels. 
    counts := make(map[string]int)
	    // bufio.NewScanner - This object can be used to read input from a specified source
	                      // os.Stdin - refers to the standard input stream, which is the default input source for a Go program.
    input := bufio.NewScanner(os.Stdin)

    //***$>_
    fmt.Println("Enter lines of text (Ctrl+D or Ctrl+Z to stop):")
    for input.Scan() {
        counts[input.Text()]++
    
    
    //***$>_
        fmt.Println("\nCount of each line:")
    // Print lines that appear more than once
        for line, n := range counts {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

