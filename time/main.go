package main

import (
	"fmt"
	"os"

	"github.com/araddon/dateparse"
)

func main() {

	timestamps := []string{"Dec 29, 2014 06:00:00 GMT", "2012-12-08", "2018-05-30"}
	for _, v := range timestamps {
		t, err := dateparse.ParseAny(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s = %v\n", v, t)
	}
}
