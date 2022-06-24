package main

import (
	"bufio"
	"fmt"
	"os"
)

//写文件
func main() {

	//OpenFile函数是一个更广泛使用的函数
	//
	path := "./test.txt"
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("file open err = %v", err)
		return
	}
	defer file.Close()

	str := "Hello,Carlos \n"
	write := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	write.Flush()
}
