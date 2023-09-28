package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main () {
	resp, err := http.Get("http://example.com/")	// Makes an HTTP GET request
	if err != nil {
		fmt.Println("!ERROR line#10 resp := http.get",err)
		return
	}
	fmt.Println(resp,"\n")				
	body, _ := ioutil.ReadAll(resp.Body)		// read the body from the response
	fmt.Println(string(body))
	resp.Body.Close()				// I had clousing the connection
}
