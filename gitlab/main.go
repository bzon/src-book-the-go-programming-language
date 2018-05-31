package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	git, err := gitlab.NewBasicAuthClient(nil, "http://localhost:10080", "root", "123qwe123")
	if err != nil {
		log.Fatal(err)
	}
	_, err = git.Groups.DeleteGroup("UnknownGroup")
	log.Println("Got here!")
	if err != nil {
		log.Fatal(err)
	}
}
