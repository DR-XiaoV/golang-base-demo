package main

func main(){
	const name string = "李欢欢"
	var   name1 string = "李欢欢"
	name1 = "韩梅梅"
	//name = "韩梅梅"  常量值不能被改变
	println(name,name1)

	// 常量在定义的时候必须被赋值

	// 多个常量一起声明
	const (
		pi float32 = 3.1415926535898
		e = 2.7182
	)
	// 同时声明多个常量  省略 就是和上一行的值相同
	const (
		baz = 100
		baz1
		baz2
	)
	print(baz1)
}