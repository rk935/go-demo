package main

import "fmt"

func recv(c chan string) {
	ret := <-c
	fmt.Println("recv: ", ret)
}

func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	//ch := make(chan string)
	//go recv(ch)
	//ch <- "王子"
	//fmt.Println("发送成功")

	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25
}
