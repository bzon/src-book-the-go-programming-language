package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	data, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	req, err := http.NewRequest("PUT", "http://localhost:8081/repository/site/go_upload/index.html", data)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "admin123")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)
	bs, err := ioutil.ReadAll(resp.Body)
	log.Println(string(bs))
}
