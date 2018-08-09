package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	dIndex := 18888
	serverIp := "127.0.0.1"
	var requestTime uint64 = 1234567890123
	ac := "dadian"

	//data := `{"ac":"migame_s_payment","app_id":"20005","create_time":1530535401,"data":{"orderId":"20180702204321_1_20799473","orderTime":1530535401,"orderType":1,"payPrice":600,"price":600,"productCode":"com.kunpo.loner.mi","productCount":1,"productName":"一米购漫画消费","uuid":1103597394},"type":"CreateOrder"}`
	data := `{"ac":"migame_s_payment","app_id":"20005","create_time":1530528372,"data":{"orderId":"","orderTime":1530528372,"orderType":0,"refundReason":"","status":0,"systemOrderId":"","uuid":0},"type":"UpdateRefund"}`

	dataBase64 := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("base64 is %v\n", dataBase64)
	encodeResult := url.QueryEscape(dataBase64)


	msg := fmt.Sprintf("ac=%s&dProcTag=%s&dIndex=%d&fromApp=%s&data=%s&ts=%d", ac, serverIp, dIndex, "milink", encodeResult, requestTime*1000)

	fmt.Printf("resutl == %v\n", msg)
	fmt.Printf("resutl == %v\n", "%2f")
}
