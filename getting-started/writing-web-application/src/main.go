package main

import (
	"html/template"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

/*** refactoring ***/
func renderTemplate(w http.ResponseWrite, tmpl string, p *Page){
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*** address persistent storage ***/
func (p *Page) save() error {		// *Page this mean that is a type of date, see line #10
					// save() is the name of the func
					// error this mean bacically that the func don't return any except and error if occour typeof error
	filename := p.Title + ".txt"		
	return os.WriteFile(filename, p.Body, 0600)		// os: documentation: https://pkg.go.dev/os#pkg-index
								// .WriteFile(): documentation: https://pkg.go.dev/os#example-WriteFile
								//exam: func WriteFile(name string, data []byte, perm FileMode) error
								//how you can see this func return a error you must handle this in the future
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
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w,"view",p)
}

/*** editHandler ***/
func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil{
		p = &Page{title:title}
	}
  renderTemplate(w, "edit", p)
}
/*** saveHandler ***/
func saveHandler(w http.ResponseWriter, r *http.Resquest){
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{title: title, Body: []byte(body)}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	http.redirect(w, r, "/view/" +title, http.StatusFound)
}

/** use */
func main(){
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}


