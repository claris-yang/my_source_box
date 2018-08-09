package mysql

import (
	"fmt"
	"util/errs"
)

type MyResult struct {
	Name string
}

type InParam struct {
	Param string
}

func CallerMysql(in InParam) (*MyResult, *errs.Error) {

	fmt.Printf("in is %v\n", in.Param)

	myResult := &MyResult{Name: "beijing"}
	fmt.Printf("tiem duration is %v\n", myResult)

	return myResult, errs.New(1002, "CallerMysql Fail, tianjin")
}
