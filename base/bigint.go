package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 大数据处理
	//设置一个大于int64的数
	a := new(big.Int)
	a, ok := a.SetString("9122322238215458478512545454878168716584545412154785452142499999000000", 10)
	if !ok {
		panic("error")
	}
	//String方法可以转换成字符串输出
	fmt.Println(a.String())

	//大数相加
	b := big.NewInt(9122322238215458)
	b = b.Mul(a, b) //  Mod 取模、Add 加、Sub 减、Mul 乘、Div 除

	fmt.Println("===b: ", b.String())
}
