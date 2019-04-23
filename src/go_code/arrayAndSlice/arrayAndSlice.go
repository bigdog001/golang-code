package main

import (
	"fmt"
)

func arrayDemo1(){
	// 定义一个数组,2种方法初始化,如果不写数组大小3那就是切片了
	var arr [3]float64 = [3]float64{3.0,4.0,5.0}
	arr[0]=1.0
	arr[1]=2.0
	arr[2]=3.0

	var total float64
	for i:=0;i<len(arr);i++{
		total+=arr[i]
	}
	average := fmt.Sprintf("%.2f", total/3)
	fmt.Println("和是",total,"平均值是",average)
}

func arrayDemo2(){
	// 第3种初始化数组的方法:
	var arr = [...]float64{3.0,4.0,5.0,78.9}
	//var arr1 = [4]float64{3.0,4.0,5.0,78.9}
	fmt.Printf("数组的地址是%p,第一个元素的地址是%v",&arr,&arr[0])

	// 第4种初始化数组的方法:指定下标
	var arr2 = [...]float64{1:3.0,2:4.0,3:5.0,0:78.9}
	fmt.Printf("数组的是%v",arr2)

	// 第5种初始化数组的方法:类型推导
	strArr := [...]string{"xiang","shou","sheng"}
	fmt.Printf("字符数组的是%v",strArr)


}
func iterateArr(){
	strArr := [...]string{"xiang","shou","sheng"}
	for i,v := range strArr{
		fmt.Printf("i=%v,v=%v \n",i,v)
	}
}
func main(){
	// 数组有值传递和地址传递,值传递的话函数执行前后变量的值互相不影响. 
	// arrayDemo1()
	// arrayDemo2()

	// 数组独有的遍历(不同于普通的遍历)
	iterateArr()
}