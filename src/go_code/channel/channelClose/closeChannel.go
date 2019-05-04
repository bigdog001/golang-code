package main

import(
	"fmt"
)

func closeChannel(){
	intchan := make(chan int,3)
	intchan <- 10
	intchan <- 20
	close(intchan)
	// 不能再写了,只能读
	num := <- intchan
	fmt.Println("",num)
}

func traverseChannel(){
	intchan := make(chan int,100)
	for i:=0;i<100;i++{
		intchan <- i*2
	}
	// 遍历时要先关闭,否则会报死锁的错误
	close(intchan)
	for v := range intchan{
		fmt.Print("\t",v)
	}
	
}

func main(){
	// closeChannel()

	// 遍历
	traverseChannel()
	
}