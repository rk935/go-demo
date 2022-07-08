package main

import "fmt"

func main() {
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[:3]
	//fmt.Println("a:", a)
	//fmt.Println("s:", s)
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	//var citySlice []string
	// 追加一个元素
	//citySlice = append(citySlice, "北京")
	// 追加多个元素
	//citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	//a := []string{"成都", "重庆"}
	//citySlice = append(citySlice, a...)
	//fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]

	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(byteS1)

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(runeS2)
}
