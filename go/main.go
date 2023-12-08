package main

import (
	"fmt"
	"sdk/go-core/modules/go-core/env"
)

func main() {

	r := env.NewEnvReaderOrExit()
	testVar := r.String("TEST_VAR", "default", "usage message")

	fmt.Println(testVar)

}
