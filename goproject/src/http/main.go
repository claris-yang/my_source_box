
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

)

func Post(url string, arg string) ([]byte, error) {

	request, err := http.NewRequest("POST", url, strings.NewReader(arg))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("charset", "UTF-8")
	request.Header.Set("END-USER-IP", "10.68.1.30")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func main() {
	
	arg := `clientInfo={ "deviceInfo": { "model": "MI 3", "miuiVersion": "4.7.1", "androidVersion": "16", "screenDensity": "320", "screenWidth": 720, "screenHeight": 1280, "marketVersionCode": 0, "isInter": false, "devicetype": 0 }, "userInfo": { "imei": "71ac0b8ad739b0e1f7447234ad1847d2", "userid": "1111", "language": "zh", "country": "CN", "serviceProvider": "WIFI", "ip": "127.0.0.1", "triggerId": "yED6Fvi2JLdH2pmx" }, "mediaType": "GAME", "contextData": { "data": { "debugv3": "on" } }, "impRequests": [ { "tagId": "1.12.4.1", "adsCount": 1, "context": { } } ] }&v=3.0&isbase64=false&appKey=GAME&sign=skipSign`
	url := "http://test.ad.xiaomi.com/post/v3"

	context, err  := Post(url, arg)
	if nil != err {	
		fmt.Printf("http post is %v\n", err)
	}

	fmt.Printf("resutl is %v\n", string(context))
}

