package main
import (
	"fmt"
	"net/http"
	"os"
	"time"
)
func main() {
	// set start time run func
	start := time.Now()
	ch0 := make(chan string)
	
	// let read what the user pass through param
	for _, url := range os.Args[1:] {
		// we can start a goroutine for use multi core
		go fetch(url,ch0)
	}
	// error control
	for range os.Args[1:] {
		fmt.Println(<-ch0) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch0 chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch0 <- fmt.Sprint(err) //sent to channel ch 0
		return
	}
	// Persistent storage
	// 0600 -> 0(special) regular file, 6(owner) has read and write permissions, 0(group), 0(otherS)
	os.WriteFile("data.txt",resp.Body,0600)
	resp.Body.Close() // don't leak resources
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}

