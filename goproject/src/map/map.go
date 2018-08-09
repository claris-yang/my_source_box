
package main

import (
	"time"
	"fmt"
)

func main() {
	m := make(map[string]interface{})


	m["nihao"] = true
	for i := 0 ; i < 1000000; i++ {
		go func() {
			for j := 0 ; j < 10000000; j++ {
				fmt.Println(m["nihao"])
			}
		}()
	}

	time.Sleep(time.Minute * 10)
}
