package main

/**
 * 通过文件加载IP信息
 *
 * @author yangtao
 * @version V1.0 创建时间：2016-11-19
 *
 */
import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type IPInfoSt struct {
	StartIP    uint32
	EndIP      uint32
	StartIPStr string
	EndIPStr   string
	Country    string
	Province   string
	City       string
}

var IpInfos []IPInfoSt

const (
	IP_FILE_DATA_ITEM_NUM = 7
)

func LoadIPInfo(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		lineData, _, err := rd.ReadLine()

		if err != nil || io.EOF == err {
			break
		}

		ipData := strings.Split(string(lineData), ",")

		if IP_FILE_DATA_ITEM_NUM > len(ipData) {
			continue
		}

		startIp, e := strconv.ParseInt(ipData[0], 10, 32)
		if nil != e {
			continue
		}

		endIp, e := strconv.ParseInt(ipData[1], 10, 32)
		if nil != e {
			continue
		}

		ipInfo := IPInfoSt{
			StartIP:    uint32(startIp),
			EndIP:      uint32(endIp),
			StartIPStr: ipData[2],
			EndIPStr:   ipData[3],
			Country:    ipData[4],
			Province:   ipData[5],
			City:       ipData[6],
		}

		IpInfos = append(IpInfos, ipInfo)
	}

/*
	for i := 0  ; i  < len(IpInfos) ; i++ {
		fmt.Println(IpInfos[i])
	}
*/


}

func main()  {
	LoadIPInfo("mi_china.txt")

	pos := BinarySearchIP(16777472)
	
	fmt.Println(pos)

	pos = BinarySearchIP(16842752)

	fmt.Println(pos)
}


func BinarySearchIP(ip uint32) (pos int) {
	start := 0
	end := len(IpInfos)

	for start < end {
		mid := start + (end-start)/2
		fmt.Println(start)
		if IpInfos[mid].StartIP >= ip {
			end = mid 
		} else {
			start = mid + 1
		}
	}

	return start
}
