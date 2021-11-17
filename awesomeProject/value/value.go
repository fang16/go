package main

import "fmt"

func changeint(a int)   {
	a=100
}

func changearray(b []int) {
	b[0]=100
}


//值类型与引用类型
//将值赋值给一个变量，并修改，若对就变量有影响则为值类型，反之为引用类型

func main()  {
	//值类型：int float32 float64 bool array 指针

	//应用：slice map
	num := 1
	changeint(num)
	fmt.Println(num)

	nums := []int{1,2,3,4}
	changearray(nums)
	fmt.Println(nums)

}
