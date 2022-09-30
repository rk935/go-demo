package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n2 uint32 = 0xFFFF
	fmt.Printf("n2 的类型 %T, 值: %d, n2占中的字节数是 %d", n2, n2, unsafe.Sizeof(n2))
}
