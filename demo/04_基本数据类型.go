package main

import  (
	"fmt"
	"math"
)

func main() {
	// 基本数据类型
	/*
	类型	描述
	uint8	无符号 8位整型 (0 到 255)
	uint16	无符号 16位整型 (0 到 65535)
	uint32	无符号 32位整型 (0 到 4294967295)
	uint64	无符号 64位整型 (0 到 18446744073709551615)
	int8	有符号 8位整型 (-128 到 127)
	int16	有符号 16位整型 (-32768 到 32767)
	int32	有符号 32位整型 (-2147483648 到 2147483647)
	int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
	*/

	var num1 int8 = 100
	fmt.Println(num1)

	fmt.Printf("%f\n",math.Pi)
	fmt.Printf("%.2f\n",math.Pi) // 保留小数点后两位

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	var flag bool = true
	if(flag){
		println("flag is:",flag)
	}

	
}