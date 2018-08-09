
package main

import (
	"reflect"
	"fmt"
)

type Hello struct {
	Name string
}

func THello( m Hello) string {
	return "hello beijing"
}

func rString(f interface{}, s... interface{}) []interface{} {

	if m, ok := f.(reflect.Method) ;ok {
		fmt.Printf("method is %v\n", m.Name)
	} else {
		fmt.Printf("f type is %v\n", reflect.TypeOf(f).NumMethod())
		fmt.Printf("f type is %v\n", reflect.ValueOf(f).String())
		fmt.Printf("f type is %v\n", reflect.ValueOf(f).NumMethod())
	}
	fc := reflect.ValueOf(f)
	in := make([]reflect.Value, len(s))
	for k, param := range s{
		in[k] = reflect.ValueOf(param)	
	}	
	
	fmt.Printf("funcName is %v\n", reflect.TypeOf(f).String())
	fmt.Println(fc.String())
	re := fc.Call(in)


	fmt.Println(re)
	var result []interface{}

/*
	for _, mi := range re {
		result = append(result, mi.interface())
	}

*/
	return result

}

type TStudent struct {
	name string
}

func TRefect(in interface{})  {
	
	var x interface{}
	
	stu := &TStudent {
		name : "tiananmen",
	}

	x = stu

	
	dst := reflect.ValueOf(&in)	
	src := reflect.ValueOf(x)
	dst.Elem().Set(src.Elem())
	in = dst.Interface()
	fmt.Printf("lalala is %v\n", in)

	v, ok := in.(*TStudent)
	if ok {
		fmt.Printf("ok ok ok is %v\n", v)
	}
}

func main() {

	m := Hello {
		Name : "beijing",
	}
	rString(THello, m)

	THello(m)

	
	stu := TStudent{name:"heibeijing"}
	TRefect(&stu)

	fmt.Printf("student is %v\n", stu)

}
