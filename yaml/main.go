package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
xxx
b: a
c: a
`

func main() {
	var p map[string]interface{}
	err := yaml.Unmarshal([]byte(data), &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got here")
	fmt.Println(string(p))
}
