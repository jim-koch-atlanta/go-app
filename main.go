package main

import (
	"fmt"

	complex "github.com/thebigkoch/go-lib/libs/complex"
)

func main() {
	fmt.Printf("Hello world!\n")
	result := complex.Times(1, 2)
	fmt.Printf("1 * 2 = %d", result)
}
