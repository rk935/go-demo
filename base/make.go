package main

import "fmt"

func main() {
	//num := 1
	//var p *int
	//
	//fmt.Println(p)
	////fmt.Println(*p)
	//
	//p = &num
	//fmt.Println("num变量的地址为：", &num)
	//fmt.Println("num变量的地址为：", p)
	//fmt.Println("指针变量p的地址为：", &p)
	//fmt.Println("指针变量p的值为：", *p)

	//data := make(map[string]int, 10)
	//var data map[string]int

	//a := make([]int, 3, 2)
	//a[0] = 0
	//a[1] = 1
	//a[2] = 2
	//fmt.Println(a)      //[0 0]
	//fmt.Println(len(a)) //2
	//fmt.Println(cap(a)) //10
	scoreMap := make(map[string]int, 1)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
}
