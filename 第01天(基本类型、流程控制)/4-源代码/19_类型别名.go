package main //必须有一个main包

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64

	var a bigint // 等价于var a int64
	fmt.Printf("a type is %T\n", a)

	type (
		long bigint
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d 类型:%T, ch = %c 类型:%T\n", b, b, ch, ch)
}
