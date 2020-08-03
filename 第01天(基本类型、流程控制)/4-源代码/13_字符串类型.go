package main //必须有一个main包

import "fmt"

func main() {
	var str1 string //声明变量
	str1 = "abc"
	fmt.Println("strl = ", str1)

	//自动推导类型
	str2 := "mike"
	fmt.Printf("str2 类型是 %T\n", str2)

	//内建函数，len()可以测字符串的长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))

	str3 := "中国nihao"
	fmt.Println("len(str3) = ", len(str3))
	fmt.Println("rune(str3) = ", rune(str3))
}
