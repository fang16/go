package main
//递归
import (
	"fmt"
)

func addN(a int) (int)  {
	if a == 1 {
		return 1
	} else {
		return a+addN(a-1)
	}
}

//阶乘

func addF( b int) (int)  {
	if b == 1 {
		return 1
	} else {
		return b * addF(b-1)
	}
}
//汉诺塔游戏
func addG(a,b,c string,N int)  {
	if N == 1 {
		fmt.Println("a","移动到","c")
		return
	} else  {
		fmt.Println("a","移动到","b")
		 addG(a,b,c,N-1)
		fmt.Println("b","移动到","c")
		addG(b,c,a,N-1)
		return

	}
}

//定义一个韩素类型的变量
func add(a,b int) int {
	return a+b
}

//callbak 格式化，将传递的数据按照每行打印还是按照一行按｜分割打印
func print(callbak func(...string),args ...string)  {
		fmt.Println("print 函数输出")
		callbak(args...)
}
func list(args... string)  {
	for i, v := range args {
		fmt.Println(i,":",v)
	}
}


func main()  {
	//函数调用
	fmt.Println(addN(9))
	fmt.Println(addF(5))
	addG("A","B","C",5)
	//定义一个函数类型的变量
	var e func(int,int) int = add
	fmt.Printf("%T\n",e)
	//callbak调用，callbak格式化函数，将传递的数据按指定格式打印
	list("A","B","C","D","E","T")

	// 匿名函数
	saihello := func(name string )  {
		fmt.Println("hello",name)
	}
	saihello("小芳")

	//定义匿名函数并调用
	func(name string) {
		fmt.Println("hi",name)
	} ("kk")

    //定义只使用一次的匿名函数
	values := func(args ...string) {
		for _,v := range args {
			fmt.Println(v)
		}
	}
	print(values,"A","B","C")

	print(func(args ...string) {
		for _,v := range args {
			fmt.Println(v,"\t")
		}
		fmt.Println()
	},"A","B","C","E")

	//闭包
	addBase := func(base int) func(int) int {
		return func(n int ) int {
			return base + n
		}
	}
	addBase4 := addBase(10)
fmt.Println( addBase4(4))
	fmt.Println(addBase(4)(4))
}
