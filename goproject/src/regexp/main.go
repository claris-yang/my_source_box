
package main

import (
	"regexp"
	"fmt"
)

func main() {
	re1 := regexp.MustCompile("[a-zA-Z]+")
	fmt.Printf("%v\n", re1.FindString("awd."))
}
