package main
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handlerFunc()
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) {
					  // r.url.path extract the path /hello from the request URL
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}


