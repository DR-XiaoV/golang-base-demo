package main

func main() {
	// iota _ 占位
	const (
		n1 = iota // 1
		n2
		n3
		_
		n4
	)
	print(n1, " ", n2, " ", n3, " ", n4, " ") // 0,1,2,4

	// iota 插队
	const (
		num1 = iota
		num2 = 100
		num3 = iota
		num4
	)
	print(num1,num2,num3,num4) // 0,100,1,2
}
