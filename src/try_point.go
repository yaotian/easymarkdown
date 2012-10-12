package main

import "fmt"


func main() {
	var p *int	   //定义一个指针
	fmt.Println(p)  //所有新定义的变量都被赋值为其类型的零值

	var i int = 9
	var p_i = &i    //获得指针的地址

	fmt.Println(p_i) 

	fmt.Println(*p_i) //获得指针指向的值

}