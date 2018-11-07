
package main

import "encoding/json"
import "fmt"

type Hello struct {
	Beijing string "json:`Beijing`"
	Riben   string "json:`Beijing`"
	IsGood  bool   `json:"isGood"`
	GameId  float64 `json:"gameId"`
	
}

func main() {
	str := `{"Beijing":"nihao","gameId":6.2234016e+07}`

	var hello Hello
	e := json.Unmarshal([]byte(str),&hello) 
	if nil != e {
		fmt.Printf("beijig is %v\n", e)
		return 
	}
	fmt.Printf("beijing is %v\n", hello)
}
