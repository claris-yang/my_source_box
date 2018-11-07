
package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%v\n", x)

	y := time.Now().Format("20060102")
	fmt.Printf("%v\n", y)

	if len(y) >= 2{
		
		fmt.Printf("%v\n", y[len(y) - 2:] )
	}

	fmt.Printf("%v", time.Now().UnixNano()/1e6)
}

