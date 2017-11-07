package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page)save() error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600)
}

func load(title string) (*Page, error) {
	f := title + ".txt"
	body, err:= ioutil.ReadFile(f)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func view(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path[len("/text/"):]
	p, _ :=load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	p := &Page{Title: "Test", Body: []byte("Welcome to the Test Page!")}
	p.save()
	http.HandleFunc("/test/", view)
	http.ListenAndServe(":8000", nil)
}
