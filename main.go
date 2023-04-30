package main

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

func handlerMain(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "header")
}

func renderTemplate(w http.ResponseWriter, tmplName string) {
	templateCache, err := createTemplateCache()

	if err != nil {
		panic(err)
	}

	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, nil)
	buffer.WriteTo(w)
}

const port = ":8080"

func main() {
	http.HandleFunc("/", handlerMain)
	
	fmt.Println("Server start on port :8080")
	fmt.Println("Please follow this link http://localhost:8080")
	http.ListenAndServe(port, nil)
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, nil
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}