package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
}

func ReadEnvValue(k string) string {
	return os.Getenv(k)
}
