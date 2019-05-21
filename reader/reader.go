package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Reader!")
	b := make([]byte, 8)
	//每个字母的ASCII码
	for {
		n, error := r.Read(b)
		//Read 用数据填充指定的字节 slice，并且返回填充的字节数和错误信息。 在遇到数据流结尾时，返回 io.EOF 错误。
		fmt.Printf("n = %v , error = %v , b = %v\n", n, error, b)
		fmt.Printf("b[:n] = %q \n", b[:n])
		if error == io.EOF {
			break
		}
	}
	/*
		n = 8 , error = <nil> , b = [72 101 108 108 111 32 82 101]
		b[:n] = "Hello Re"
		n = 5 , error = <nil> , b = [97 100 101 114 33 32 82 101]
		b[:n] = "ader!"
		n = 0 , error = EOF , b = [97 100 101 114 33 32 82 101]
		b[:n] = ""
	*/
}
