package main

import "flag"

type flagOpts struct {
	value *string
}

var flagVals flagOpts
var value = flag.String("value", "", "some value")

func main() {
	flag.Parse()
	flagVals.value = value
}
