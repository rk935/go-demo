package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zzn99/demo/address"
)

func main() {
	obj := &address.Person{
		Id:     1,
		Name:   "zhaonan",
		Gender: address.GenderType_FEMALE,
		Number: "123123",
	}
	data, _ := proto.Marshal(obj)

	fmt.Println("data: ", string(data))

	unObj := &address.Person{}
	proto.Unmarshal(data, unObj)
	fmt.Println(unObj)
}
