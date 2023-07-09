package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func main () {
	counts := make(map[string]int)
														// take 2th params
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		fmt.Println(data)
		binary := ""
		for i := 0; i < len(data); i++ {
			binary += strconv.FormatInt(int64(data[i]),2)
		}
		fmt.Println(binary)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 : %v\n", err)
			continue
		}
												// string.Split - dividir una cadena (string) en una lista de substrings basándose en un separador específico	
		for _, line := range strings.Split(string(data), "\n"){

			counts[line]++
			fmt.Println(counts)
		}
	}
	for line, n := range counts {
		fmt.Printf(line)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
