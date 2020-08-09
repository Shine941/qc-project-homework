package main

import (
	"fmt"

	"github.com/01/hello/unity"
)

func main() {
	var ty int
	var x string
	fmt.Println("请输入h或其他")
	fmt.Scanln(&x)
	unity.Choose(x)
	fmt.Println("请输入type的值：0或1")
	fmt.Scanln(&ty)
	if ty == 1 {
		fmt.Println("春眠不觉晓")
	} else if ty == 0 {
		fmt.Println("似花还似非花,也无人惜从教坠")
	} else {
		fmt.Println("输入错误")
	}
}
