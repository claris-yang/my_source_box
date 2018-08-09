
package main

import (
	"fmt"
	"sync"
)


var m sync.Mutex

func main() {
	wg := new(sync.WaitGroup)

	value := make(map[string]bool)

	for i := 0 ; i < 100000; i++ {
		wg.Add(1)
		go func() {

			m.Lock()
			defer m.Unlock()

			if v, ok := value["nihao"] ; !ok {
				fmt.Printf("map key nihao not exist!")
			} else {
				fmt.Printf("map value is %v\n", v)
			}

			wg.Done()
		}()
	}
	for i := 0 ;i < 100000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			value["nihao"]  = true

			wg.Done()
		}()
	}

	wg.Wait()
}
