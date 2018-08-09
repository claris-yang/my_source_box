package controller

import (
	"cms-util/tracing"
	"controller/mysql"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

func Serve() {
	regHandler()
	beego.Run()
}
func regHandler() {
	const (
		//zipkinUrl = "http://localhost/zipkin/api/v1/spans"
		//zipkinUrl = "http://dev.knights.g.mi.srv/zipkin/api/v1/spans"
		zipkinUrl = "http://dev.knights.g.mi.srv/api/v1/spans"
	)
	tracing.Init(zipkinUrl, "test-monitor")
	beego.Router("/hello", &HelloController{}, "get,post:Get")
	fmt.Printf("tiancai\n")

}

type HelloController struct {
	tracing.TraceBeegoController
}

func (this *HelloController) Get() {
	this.StartTrace()
	time.Sleep(time.Second)
	fmt.Printf("header is %v\n", this.Ctx.Request.Header.Get("open-trace"))
	fmt.Printf("header is %v\n", this.Ctx.Request.Header.Get("Accept"))
	fmt.Printf("header is %v\n", this.Ctx.Request.Header.Get("Host"))

	input := mysql.InParam{Param: "xinjiapo"}
	result, err := this.Caller("self mysql", mysql.CallerMysql, input)
	if nil != err {
		if val, ok := result[0].(mysql.MyResult); ok {
			fmt.Printf("guangzhou is %v\n", val)
		}
	}
	defer this.FinishTrace()
	response := make(map[string]interface{})
	response["hello"] = "beijing"
	this.Data["json"] = response
	this.ServeJSON()
}
