
package main

import (
	"fmt"
)

func main() {
	c := make( chan int, 10)
	c <- 10
	for ci := range(c) {
		fmt.Printf("ok is %v", ci)
	}

	fmt.Printf("over")
}
