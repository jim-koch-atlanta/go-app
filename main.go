package go_app

import (
	"fmt"

	complex "github.com/thebigkoch/go-lib/libs/complex"
)

func go_app() {
	fmt.Printf("Hello world!\n")
	result := complex.Times(1, 2)
	fmt.Printf("1 * 2 = %d", result)
}
