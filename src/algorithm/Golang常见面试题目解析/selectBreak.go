package main

import (
	"fmt"
	"strings"
)

func main() {
	//tick := time.NewTicker(2 * time.Second)
	//var a int
	//for {
	//	select {
	//	case <-tick.C:
	//		a++
	//		if a > 2 {
	//			break
	//		}
	//	}
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("循环还在执行")
	//}
	str := "ABC"
	i := 0
	fmt.Print(string(rune(str[i])))
	fmt.Println(strings.Count(str, ""))
}
