
package main

import (
	"strings"
	"regexp"
	"fmt"
	"runtime"
//	"runtime/debug"
	"time"
)

func opentrace() {
}

func fun1() {
	//fmt.Printf("im fun 1 %v\n", stack())
	opentrace()
	fun2()
	fmt.Printf("find father fun1 is %v\n", FindCallerFuncName("main.fun1"))
}

func fun2() {
	//fmt.Printf("im fun 2 %v\n", stack())
	go fun3()	

	fmt.Printf("find father is fun2 %v\n", FindCallerFuncName("main.fun2"))
	time.Sleep(time.Second * 2)
}

func fun3() {

	fmt.Printf("find father is fun3 %v\n", FindCallerFuncName("main.fun3"))


	go fun4()

	time.Sleep(time.Second * 2)
}

func fun4() {
//	debug.PrintStack()
	fmt.Println(stack())
}

type THello struct {
}

func (h *THello) Hello() {
	fun1()
}

func main() {
	hello := THello{}
	hello.Hello()
//	funtmp()
}


func stack() string {

	var buf [2<<10] byte
	l := runtime.Stack(buf[:], true)
	return string(buf[:l])
	
}

func funtmp() {
	getStackLink()
	
	time.Sleep(time.Second * 2)
}

func RegexpFuncName(s string) string {

	reg := regexp.MustCompile(`(created by ([[:ascii:]]*\.[_|0-9|a-z|A-Z][_|0-9|a-z|A-Z]*))|([[:ascii:]]*\.[_|a-z|A-Z][_|0-9|a-z|A-Z]*)`)
	result := reg.FindAllStringSubmatch(s, 1)

	if len(result) > 0 {
		if len(result[0]) > 2 && "" != result[0][2] {
			return result[0][2]	
		}  else if len(result[0]) > 3 && "" != result[0][3] {
			return result[0][3]
		}
	}
	return ""
	
}

func FindCallerFuncName(selfFunName string) string {
	var findFlag bool = false
	stackInfo := getStackLink()
	for _, stack := range stackInfo {
		for _, funs := range stack {
			if findFlag {
				 return RegexpFuncName(funs)
			
			}
			if strings.Contains(funs, selfFunName)  {
				findFlag = true
			}
		}
	}

	return ""
}


func getStackLink() [][]string {

	var stackLink [][]string
	var buf[2 << 16]byte
	l := runtime.Stack(buf[:], true)
	if l == len(buf) {
		fmt.Printf("stack buf is too small\n")
	}

	reg := regexp.MustCompile(`(goroutine[[:ascii:]]*\n{2}?)|(goroutine[[:ascii:]]*\n{1}?)`)

	result := reg.FindAllString(string(buf[:l]), - 1)

	for _, val := range result {	
		var funarr []string
	
		gfuns := strings.Split(val, "\n")
		
		for idx, f := range gfuns {
			if 1 == idx % 2 {
				funarr =  append(funarr, f)	
			}
		}

		stackLink = append(stackLink, funarr)
	}

	return stackLink
}

/*
func stack() []byte {
	buf := new(bytes.Buffer) // the returned data // 返回的数据
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	// 我们的循环打开文件并读取它们。这些变量记录了当前已加载的文件。
	var lines [][]byte
	var lastFile string
	// 我们关心 Caller 的调用者，因此跳过2帧
	for i := 2; ; i++ { // Caller we care about is the user, 2 frames up
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		// 至少要打印这么多信息。如果我们找不到来源，它就不会被显示。
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		// 在栈跟踪中，行号从1开始，单我们的数组下标却是从0开始
		line-- // in stack trace, lines are 1-indexed but our array is 0-indexed
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}
*/

