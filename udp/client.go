package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		//IP: net.IPv4(49, 0, 206, 185),
		//IP: net.IPv4(192, 168, 56, 101),
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 12223,
	})
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	fmt.Println("upd 客户端启动...")

	defer socket.Close()

	for {
		// 发送数据
		senddata := []byte("hello server!")
		_, err = socket.Write(senddata)
		if err != nil {
			fmt.Println("发送数据失败!", err)
			return
		}
		time.Sleep(time.Second)
		// 接收数据
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			return
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n", data)
	}
}
