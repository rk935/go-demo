package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	data := map[string]string{"123": "2"}
	fmt.Printf("%v \n", data)
	fmt.Printf("%+v \n", data)
	fmt.Printf("%#v \n", data)

	var a float32 = 3.14
	var x interface{}
	x = a
	b, ok := x.(float32)
	if ok {
		fmt.Printf("类型断言成功,b类型为%T\n", b)
	} else {
		fmt.Println("类型断言不成功")
	}

	// byte转io.Reader
	byteData := []byte("hello world")
	reader := bytes.NewReader(byteData)
	fmt.Println("reader: ", reader)

	buf := make([]byte, len(byteData))
	_, err := reader.Read(buf)
	if err != nil {
		fmt.Println("buf err: ", err)
	}
	fmt.Println("buf", string(buf))

	// io.Reader 转 []byte
	ioReaderData := strings.NewReader("Hello AlwaysBeta")
	buffer := &bytes.Buffer{}
	if _, err = buffer.ReadFrom(ioReaderData); err != nil {
		fmt.Println("buf err: ", err)
	}
	bufData := buffer.Bytes()
	fmt.Println(string(bufData))

	var bs strings.Builder
	bs.WriteString(`{"title" : "`)
	bs.WriteString("liuzi")
	bs.WriteString(`"}`)
	fmt.Println("bs:", bs.String())
}
