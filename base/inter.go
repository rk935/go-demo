package main

import (
	"fmt"
)

// Sayer 接口
type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

// Dog 狗结构体类型
type Dog struct {
	Name string
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

func main() {
	//var x Mover // 声明一个Mover类型的变量x

	//var d1 = Dog{} // d1是Dog类型
	//x = d1         // 可以将d1赋值给变量x
	//x.Move()
	//
	//var d2 = &Dog{} // d2是Dog指针类型
	//x = d2          // 也可以将d2赋值给变量x
	//x.Move()
	//fmt.Println("数据类型: ", reflect.TypeOf(x))

	//var c1 = &Cat{} // c1是*Cat类型
	//x = c1          // 可以将c1当成Mover类型
	//x.Move()

	// 下面的代码无法通过编译
	//var c2 = Cat{} // c2是Cat类型
	//x = c2         // 不能将c2当成Mover类型
	//Mover(&Cat{}).Move()

	var d = Dog{Name: "旺财"}

	var s Sayer = d
	var m Mover = d

	s.Say()  // 对Sayer类型调用Say方法
	m.Move() // 对Mover类型调用Move方法
}
