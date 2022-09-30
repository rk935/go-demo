package main

import (
	"bytes"
	"fmt"
	_ "github.com/robfig/cron/v3"
)

func main() {
	s1 := "hello"
	s2 := "world"

	var bt bytes.Buffer
	bt.WriteString(s1)
	bt.WriteString(s2)
	fmt.Println("result: ", bt.String())
}
