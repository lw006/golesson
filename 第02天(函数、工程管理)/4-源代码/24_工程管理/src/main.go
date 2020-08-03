package main //必须

func main() {
	test() // main.go文件拆分后在同目录下的其他文件要想一起go run 则使用 go run . 或者使用go build 然后执行二进制文件
}
