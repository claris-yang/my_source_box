package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
	"util/maps"
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

/*      设置超时时间 */
var client = &http.Client{

	Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Duration(500000000))

			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Duration(50000000000)))

			return conn, nil
		},
		ResponseHeaderTimeout: time.Duration(50000000000),
	},
}

//var client = &http.Client{}

func Get(url string) ([]byte, error) {
	//      向服务端发送get请求

	response, err := client.Get(url)
	if nil != err {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		return str, nil
	}
	return nil, err
}

func main() {

	/*
		arg := `clientInfo={ "deviceInfo": { "model": "MI 3", "miuiVersion": "4.7.1", "androidVersion": "16", "screenDensity": "320", "screenWidth": 720, "screenHeight": 1280, "marketVersionCode": 0, "isInter": false, "devicetype": 0 }, "userInfo": { "imei": "71ac0b8ad739b0e1f7447234ad1847d2", "userid": "1111", "language": "zh", "country": "CN", "serviceProvider": "WIFI", "ip": "127.0.0.1", "triggerId": "yED6Fvi2JLdH2pmx" }, "mediaType": "GAME", "contextData": { "data": { "debugv3": "on" } }, "impRequests": [ { "tagId": "1.12.4.1", "adsCount": 1, "context": { } } ] }&v=3.0&isbase64=false&appKey=GAME&sign=skipSign`
		url := "http://test.ad.xiaomi.com/post/v3"

		context, err := Post(url, arg)
		if nil != err {
			fmt.Printf("http post is %v\n", err)
		}

		fmt.Printf("resutl is %v\n", string(context))
	*/

	result := make(map[string]int)

	//url := "http://app.knights.mi.com/knights/recommend/simple/page/normal/v5?page=1&imei=1234566&uuid=1231231&id=16233"
	url := "http://10.105.45.42/knights/recommend/simple/page/normal/v5?page=1&imei=1234566&uuid=1231231&id=16233"
	for i := 0; i < 1000; i++ {
		resp, e := Get(url)
		if e != nil {
			fmt.Printf("beijing err is %v\n", e)
			return
		}

		var data map[string]interface{}
		e = json.Unmarshal(resp, &data)
		if nil != e {
			fmt.Printf("riben is %v\n", e)
			return
		}

		small := maps.GetMap(data, "data")

		blocks := maps.GetArrMap(small, "blocks")
		for _, block := range blocks {
			id := maps.GetString(block, "id")
			if id == "16561" {
				list := maps.GetArrMap(block, "list")
				for _, small := range list {
					gameId := maps.GetString(small, "id")
					if val, ok := result[gameId]; ok {
						result[gameId] = val + 1
					} else {
						result[gameId] = 1
					}
				}
			}
		}
	}

	for k, v := range result {
		fmt.Printf("gameid = %v, nums = %v\n", k, v)
	}
}
