package main

import (
	"html/template"
	"net/http"
)

const (
	staticFilesDir = "/Users/spathak/projects/shani1998.github.io/"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	// do whatever you need to do

	//myvar := map[string]interface{}{"MyVar": "Foo Bar Baz"}
	//fmt.Println(skills)
	//data, _ := json.Marshal(skills)
	//fmt.Println(string(data))
	data := ""
	outputHTML(w, staticFilesDir+"/index.html", string(data))
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
