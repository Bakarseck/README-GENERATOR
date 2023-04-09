package main

import "os"

func main() {
	arguments := os.Args[1:]
	filename := arguments[0]
	title := arguments[1]
	description := arguments[2]
	author := arguments[3]

	textReadme := 
	`<h1 align="center">` + title + `<h1>
	<h2> Description <h2>
	<p>` + description + `</p>
	<br>
	<br>
	<br>
	<a href="https://github.com/` + author + `">` + author + `</a>`

	os.WriteFile(filename + ".md", []byte(textReadme), 0666)
}