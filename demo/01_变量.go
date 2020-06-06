package main

import (
	"fmt"
)
func main() {
	fmt.Println("基础变量声明")
	base()
}
// 声明变量 var 变量name 变量类型 = 初始值
func base(){
	var num int = 3123
	var str string = "初见"
	var flag bool = false
	var (
		num1 int
		str1 string
		flag1 bool
		num2 float32
	)
	num1 = 321
	str1 = "chujian"
	flag = false
	num2 = 43.132
	println(num,num1,num2,str,str1,flag,flag1,)

	var age,name = 123,"huawei"
	println(age,name)

	var x,_ = foo()
	var _,y = foo()
	fmt.Println(x,y)
	// _ 匿名变量 不占用命名空间 不会分配内存  匿名变量之间不会存在重复命名

	//函数以外的每个语句都必须以关键字 var const func 开头
	// := 不能在函数之外使用
	// _ 多用于占位便是忽略值
}

func foo()(int,string){
	return 10,"chujian"
}