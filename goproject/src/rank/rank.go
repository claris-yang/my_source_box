
package main

import (
	"math/rand"
	"time"
	"fmt"
)

func RandArray(list []interface{}) []interface{} {

	for idx := len(list) - 1; idx >= 0; idx-- {
		randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(idx+1) % (idx + 1)
		t := list[idx]
		list[idx] = list[randNum]
		list[randNum] = t
	}

	return list
}

func main() {

/*
	views := 4
	nums := make([] int, 4)
	var i uint64
	for i = 0 ; i < 10000; i++ {
		fmt.Println(i)
		randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(views) % views
		nums[randNum] += 1
	}

	for idx, val := range nums {
		fmt.Printf("No%d : %v\n", idx, val)
	}
*/
	list := make([]interface{} , 11)
	list[0] = 1
	list[1] = 2
	list[2] = 3
	list[3] = 4
	list[4] = 5
	list[5] = 6
	list[6] = 7
	list[7] = 8
	list[8] = 9
	list[9] = 10
	list[10] = 11
	for i := 0 ; i < 10 ; i++ {
		list = RandArray(list)
		for _, val := range list {
			fmt.Printf("%v, ", val)
		}
		fmt.Println()
	}
	
/*

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
*/
}
