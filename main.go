package main

import (
	"os"
	"text/template"
)

func main() {
	fileContent := `{{ . }}`
	t := template.Must(t.New("fileContent")).Parse()
	t.Execute(os.Stdout, "var")
}
