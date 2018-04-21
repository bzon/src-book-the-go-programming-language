package main

import (
	"log"
	"os"
	"text/template"
)

const templ = `word is = {{ .Message }}
to be sent to: {{ range .Receipients }}
- {{.}}@accenture.com{{ end }}
`

type Word struct {
	Message     string
	Receipients []string
}

func main() {
	output1, err := template.New("output1").Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	output1.Execute(os.Stdout, &Word{Message: "Hello World", Receipients: []string{"john", "jb"}})
}
