package main

import (
	"fmt"
	"os"
)

type Project struct {
	Title        string
	Description  string
	Install      string
	Usage        string
	Dependencies string
	Contributors string
}

func main() {
	if len(os.Args) < 7 {
		fmt.Println("Usage: go run main.go <title> <description> <install> <usage> <dependencies> <contributors>")
		return
	}

	project := Project{
		Title:        os.Args[1],
		Description:  os.Args[2],
		Install:      os.Args[3],
		Usage:        os.Args[4],
		Dependencies: os.Args[5],
		Contributors: os.Args[6],
	}

	file, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "# %s\n\n", project.Title)
	fmt.Fprintf(file, "%s\n\n", project.Description)
	fmt.Fprintf(file, "## Installation\n\n```console\n%s\n\n```\n\n\n", project.Install)
	fmt.Fprintf(file, "## Usage \n\n\n%s\n\n", project.Usage)
	fmt.Fprintf(file, "## Dependencies\n\n%s\n\n", project.Dependencies)
	fmt.Fprintf(file, "## Contributors\n\n%s\n", project.Contributors)

	fmt.Println("README.md created successfully.")
}
