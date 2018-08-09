
package main

import (
	"runtime"	
	"fmt"
	"time"
)


func print_stack() {
	var buf[2 << 10] byte
	l := runtime.Stack(buf[:], true)
	fmt.Printf("stack : %v\n", string(buf[:l]))
	time.Sleep(time.Second)
}

func main() {
	for i := 0 ; i < 1000000; i++ {
		go print_stack()
	}

	time.Sleep(time.Second * 60)
}
