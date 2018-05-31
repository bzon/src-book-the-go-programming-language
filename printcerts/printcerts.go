package main

import (
	"crypto/x509"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	certs, err := x509.SystemCertPool()
	end := time.Now()
	if err != nil {
		panic(err)
	}
	fmt.Printf("found %d certs in %v\n", len(certs.Subjects()), end.Sub(start))
}
