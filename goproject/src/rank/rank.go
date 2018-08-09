
package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {

	var num0 , num1, num2, num3, num4 int 
	for i := 0 ; i < 1000000; i++ {

	randAddress := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5) % 5

	
	fmt.Println(randAddress)
	if randAddress == 0 {
		num0 += 1;
	} else if randAddress == 1 {
		num1 += 1;
	} else if randAddress == 2 {
		num2 += 1;
	} else if randAddress == 3 {
		num3 += 1;
	} else if randAddress == 4 {
		num4 += 1;
	}
	}

	fmt.Println(num0)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

fmt.Printf("")
}
