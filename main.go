// main 方法必须要在 main 包里面
package main

import "fmt"

import "rsc.io/quote"

// 无参数无返回值
func main() {
	fmt.Println(quote.Go())
}

// go run main.go 就可以执行
// 如果文件不叫 main.go 则需要 go build 得到可运行的文件，
// 直接运行即可
