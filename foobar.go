package main

import (
	"fmt"

	"github.com/sdlowrey/mocktest/foo"
)

func main() {
	foos := make([]*foo.Foo, 3)

	for f := range foos {
		foos[f] = foo.New()
	}

	for f := range foos {
		fmt.Printf("%d : %v\n", f, foos[f].Content)
	}
}