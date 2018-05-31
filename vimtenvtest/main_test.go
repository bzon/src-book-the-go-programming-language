package main

import (
	"os"
	"testing"
)

var k = os.Getenv("GITLAB_USERNAME")

func TestReadEnv(t *testing.T) {
	if k == "" {
		t.Fatalf("%s is empty", k)
	}
}
