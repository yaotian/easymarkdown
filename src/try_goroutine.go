package main

import (
	"fmt"
	"time"
	"runtime"
)

var c chan int

func ready(w string, sec int) {	
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w,"is ready")
	c<-1
}

func main() {
	runtime.GOMAXPROCS(2)  //因为有两个cpu
	c = make(chan int)  

	go ready("Tea", 1)
	go ready("Coffee", 10)	
	fmt.Println("I'm waiting")
	time.Sleep(5 * time.Second)   // Coffee 不会出现，因为这里只等了5秒，程序退出，所有的goroutine都会退出


	var i int = 0
	L:	for{
		select {
		case <-c:
			i++
			if i>1{
				break L
			}
		}
	}
	fmt.Println("completed!")
}

