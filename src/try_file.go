package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_whole() {
	buf := make([]byte, 1024)
	f, _ := os.Open("test.md") //打开文件，os.Open 返回一个实现了io.Reader 和io.Writer 的*os.File；
	defer f.Close()
	for {
		n, _ := f.Read(buf) //一次读取1024字节
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}

}

func read_line() {
	f, _ := os.Open("test.md")
	defer f.Close()
	r := bufio.NewReader(f)
	s, ok := r.ReadString('\n')
	if ok != nil {
		fmt.Println(s)
	}

}
func main() {
	read_line()
}
