
package main

import (
	"time"
	"fmt"
)

func testGo() {
	for i := 0 ; i < 10; i++ {
		go func(idx int) {
			for {
				fmt.Println(idx);
			}
		}(i)
	}
}

func testCh(iCh chan int) {
	
	for i := range iCh  {
		
		fmt.Printf("%v \n", i )	
	}	
	fmt.Println(1)
/*
	n, m := <- iCh
	fmt.Printf("%v, %v\n", n, m)
*/
}

func main() {
	ch := make(chan int)
	go testCh(ch)
	ch <- 1	
	ch <- 2	
	ch <- 3	
	ch <- 4	
	ch <- 5	
	ch <- 6	
	close(ch)
	time.Sleep(time.Second * 10)

}
