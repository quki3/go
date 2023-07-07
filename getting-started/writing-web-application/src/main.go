package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

/*** address persistent storage ***/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/*** load Pages ***/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename) // The _ (underscore) symbol is used to ignore the second return value, which is an error indicating whether the file reading operation encountered any errors.
	if err != nil {
		return nil,err
	}
	return &Page{Title:title, Body:body},nil
}

/*** view Handler ***/
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

/** use */
func main(){
	http.HandleFunc("/view/",viewHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}


