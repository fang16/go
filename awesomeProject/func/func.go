package main

import "fmt"

//定义函数,无参无返回值
//
func sayHelloword() {
	fmt.Println("hello word!")
}

//定义有参函数（形参）
func sayHi(name string)  {
	fmt.Println("你好：",name)
}
//第三个int定义返回值类型
func add(a int,b int,c int) int  {
	return a+b+c
}

//函数参数的类型合并//
func and(a,b,c int) int {
	return a+b+c
}

func readd(a,b int) (int,int,int,int)  {
	return a+b,a-b,a*b,a/b

}
//函数的可变参数,args 是一个切片类型
func adnN(a,b int,args ...int) int {
	aa := a+b
	for _ , v := range args {
		aa += v
	}
return aa
}

//命名返回值
func desc1(a,b int) (dd int,ee int,cc int,ss int )  {
	dd = a+b
	ee = a-b
	cc = a*b
	ss = a/b
	return
}

//可变参数的解包,args 切片类型，args...切片解包,切片解包只能在函数中使用
func cacl(op string,a,b int , args ...int) int  {
	switch op  {
	case "add" :
		return adnN(a,b , args...)
	}
	return -1

}

func main()  {
	fmt.Printf("%T\n",sayHelloword)

	//调用函数（实参）
	sayHi("小芳")
	a := add(3,3,3)
	fmt.Println(a)
	b := and(3,3,3)
	fmt.Println(b)
	c := adnN(1,2,3,3,4,5)
	fmt.Println(c)
	fmt.Println(cacl("add",1,2,4,5,6))
	args := []int{1,2,3,4}
	fmt.Println(cacl("add",1,2,args...))

	//append 移除应用,h...为解包操作
	f := []int{1,2,3,4}
	g := f[1:]
	h := f[:1]
	fmt.Println(g,h)
	fmt.Println(append(g,h...))
	i,j,k ,l := readd(9,3)
	fmt.Println(i,j,k,l)

	fmt.Println(desc1(3,3))


}

