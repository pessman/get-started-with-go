package main

import (
	"fmt"

	"github.com/pessman/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Ashe"))
}
