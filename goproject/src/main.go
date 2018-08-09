package main

import (
	"RxGo/handlers"
	"RxGo/observable"
	"cms-util/monitor"
	"cms-util/tracing"
	"controller"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ConvNum(ori, c int) bool {
	var y int = ori % c
	for {

		newY := ori % c
		if y != newY {
			return false
		}

		tmp := ori / c
		y = newY
		ori = tmp

		if tmp == 0 {
			break
		}
	}
	return true

}
func solution(line string) string {
	num, e := strconv.Atoi(line)
	if nil != e {
		return "0"
	}

	for i := 2; i <= num; i++ {
		flag := ConvNum(num, i)
		if flag {
			return strconv.Itoa(i)
		}
	}
	return "0"
}

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] > c[j]
}

func solution1(line string) string {
	// 在此处理单行数据
	input := strings.Split(line, " ")
	var nums IntSlice
	for _, val := range input {
		num, e := strconv.Atoi(val)
		if nil != e {
			continue
		}

		nums = append(nums, num)
	}

	sort.Sort(nums)

	var pri int = 0

	if len(nums) > 0 {
		pri = nums[0]
	}

	var sum int = 0
	for _, val := range nums {
		if pri-1 == val || pri == val {
			sum += val
		}
	}

	// 返回处理后的结果
	return strconv.Itoa(sum)
}

func TestRxGo() {
	f1 := func() interface{} {
		// Simulate a blocking I/O
		time.Sleep(2 * time.Second)
		fmt.Printf("f1\n")
		return 1
	}
	f2 := func() interface{} {
		// Simulate a blocking I/O
		time.Sleep(time.Second)
		fmt.Printf("f2\n")
		return 2
	}

	onNext := handlers.NextFunc(func(v interface{}) {
		fmt.Printf("nextFunc is %v\n", v)
	})
	wait := observable.Start(f1, f2).Subscribe(onNext)
	sub := <-wait
	if err := sub.Err(); err != nil {
		fmt.Printf("heppend error")
	}
}

func TestPost() {

	priorAr := fmt.Sprintf("appKey=GAME&clientInfo=%v&v=3.0",
		`{"deviceInfo":{"model":"MI 3","miuiVersion":"4.7.1","androidVersion":"16","screenDensity":"320","screenWidth":720,"screenHeight":1280,"marketVersionCode":0,"isInter":false,"devicetype":0},"userInfo":{"imei":"71ac0b8ad739b0e1f7447234ad1847d2","userid":"1111","language":"zh","country":"CN","serviceProvider":"WIFI","ip":"127.0.0.1","triggerId":"yED6Fvi2JLdH2pmx"},"mediaType":"GAME","contextData":{"data":{"debugv3":"on"}},"impRequests":[{"tagId":"1.12.4.1","adsCount":3,"context":{}}]}`,
	)

	fmt.Printf("query is %v\n", url.QueryEscape(priorAr))

	hostName := "http://test.ad.xiaomi.com/post/v3"

	fullUrl := fmt.Sprintf("%v?%v", hostName, priorAr)
	u, err := url.Parse(fullUrl)
	if nil != err {
		return
	}

	arg := fmt.Sprintf("%v%v%v%v&appSecret=7f1526fb86c71db34b576d96e2e7d940",
		"POST\n",
		"test.ad.xiaomi.com\n",
		"/post/v3\n",
		//fullUrl
		strings.Replace(u.Query().Encode(), "+", "%20", -1),
	)

	h := md5.New()
	h.Write([]byte(arg))
	md5CheckData := h.Sum(nil)

	fmt.Printf("old=%v\n", arg)

	postArg := fmt.Sprintf("appKey=GAME&clientInfo=%v&v=3.0&sign=%v",
		`{"deviceInfo":{"model":"MI 3","miuiVersion":"4.7.1","androidVersion":"16","screenDensity":"320","screenWidth":720,"screenHeight":1280,"marketVersionCode":0,"isInter":false,"devicetype":0},"userInfo":{"imei":"71ac0b8ad739b0e1f7447234ad1847d2","userid":"1111","language":"zh","country":"CN","serviceProvider":"WIFI","ip":"127.0.0.1","triggerId":"yED6Fvi2JLdH2pmx"},"mediaType":"GAME","contextData":{"data":{"debugv3":"on"}},"impRequests":[{"tagId":"1.12.4.1","adsCount":3,"context":{}}]}`,
		hex.EncodeToString(md5CheckData),
	)

	fmt.Printf("i will post \n %v\n", postArg)

	rsp, err := Post(hostName, postArg)
	if nil != err {
		fmt.Printf("err is %v\n", err)
	}

	fmt.Printf("rsp is content %v\n", string(rsp))
}

func Post(url string, arg string) ([]byte, error) {

	request, err := http.NewRequest("POST", url, strings.NewReader(arg))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("charset", "UTF-8")
	request.Header.Set("END-USER-IP", "10.235.149.239")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func StartMonitor() {
	monitor := monitor.NewMonitor("lugu", "com.mi.knights.app")
	monitor.AddRouteMonitor("comics", "https://app.knights.mi.com/knights/contentapi/comics/info?id=90610011")
	monitor.AddRouteMonitor("another", "http://localhost:8080/hello")

}

func CallMistore() {
	time.Sleep(time.Second)
}

func StartTracing() {

	const (
		zipkinUrl = "http://localhost:9411/api/v1/spans"
	)
	tracing.Init(zipkinUrl, "test-monitor")

}

func main() {
	StartMonitor()
	controller.Serve()
}
