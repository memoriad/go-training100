package main

import (
	"fmt"
	"strconv"
	"training-go/address"
	"training-go/condition"
	"training-go/finally"
	"training-go/hashmap"
	"training-go/hello"
	"training-go/list"
	"training-go/loop"
	"training-go/structure"
)

func main() {
	const title = "hello"

	hello.Hello(title)

	printSwap()

	fmt.Println(strconv.Itoa(loop.SumLoop()))

	fmt.Println(strconv.Itoa(condition.SumIf()))

	finally.AuditLog()

	i := 20
	address.SetValue(&i)
	fmt.Println(i)

	list.PrintArray()
	list.PrintSlice()

	loop.LoopSlice()

	hashmap.GetMap()

	v1 := &structure.Vertex{
		X: 20,
		Y: 50,
	}

	fmt.Println("X:", v1.GetX())
	fmt.Println("Y:", v1.GetY())

	v1.SetX(100)
	v1.SetY(120)

	printVertexer(v1)

	v2 := &structure.Vertex2{
		X: 20,
		Y: 50,
	}

	printVertexer(v2)

	var obj interface{}
	obj = "123"
	obj = "hello"

	fmt.Println("obj:", obj.(string))
}

func printVertexer(v structure.Vertexer) {
	fmt.Println("X:", v.GetX())
	fmt.Println("Y:", v.GetY())
}

func printSwap() {
	fmt.Println(hello.Swap(1, 2))
}
