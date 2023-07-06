package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

/*** address persistent storage ***/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.Writefile(filename, p.Body, 0600)
}

/*** load Pages ***/
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename) // The _ (underscore) symbol is used to ignore the second return value, which is an error indicating whether the file reading operation encountered any errors.
	return &Page{Title:title, Body:body}
}
