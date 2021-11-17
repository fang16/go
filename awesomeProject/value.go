package main

import (
	"fmt"
)
import (
	"math"
)

const test = true
const n = 10000
func main() {
	//定义变量
	/*联系
			go
	变量定义了必须要使用
	*/
	var i, c, d int = 1, 2, 3
	fmt.Println(i, c, d)
	var e = true
	fmt.Println(e)
	var f int
	fmt.Println(f)
	g := f
	fmt.Println(g)
	//const用于声明常量
	const cl string = "changliang"
	fmt.Println(cl, test)
	fmt.Println(math.Sin(n))
	fmt.Println(math.Cos(n))

	//枚举
	const (
		E1 int = iota
		E2
		E3
		E4
	)
	const (
		E5 int8 = iota
		E6
		E7
		E8
	)
	fmt.Println(E1, E2, E3, E4)
	fmt.Println(E5, E6, E7, E8)

	//作用域:定义标识符可使用范围，使用{}来定义
	outer := "OUTER"
	fmt.Printf(outer)
	{
		vcs := "VCS"
		fmt.Println(vcs)
	}
	vcs := "VCS"
	fmt.Println(vcs)

	fmt.Println("1+1=", 1+1)
	fmt.Print("2+2=", 2+2, vcs)
	fmt.Printf("\n%T,%s,%T,%d\n", vcs, d, 1, 2)

	//for 循环
	z := 1
	for z <= 4 {
		fmt.Println(z)
		z = z + 1
	}
	for j := 4; j <= 9; j++ {
		fmt.Println(j)
	}

	//布尔类型，不可赋值
	var zero bool
	isboy := true
	isgirl := false
	fmt.Println(zero,isboy,isgirl)

	//逻辑运算 &&和  ｜｜或
	fmt.Println("&&")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(false || true)

	//!aBool 取反
	fmt.Println(!true)
	//判断
	fmt.Println(isboy == isgirl)
	fmt.Println(isboy != isgirl)
	fmt.Printf("%T",false)
	fmt.Printf("%t,%s,%d",isboy,isgirl,1)

	//整数，共有13种类型，int为有符号类型，位数根据系统走
	var zs int = 34
	fmt.Printf("%T,%d\n",zs,zs)
	fmt.Println("fen")
	fmt.Println(1+2)
	fmt.Println(2-2)
	fmt.Println(9/4)
	fmt.Println(9%4)
	fmt.Println(2*2)
	zs++
	fmt.Println(zs)
	zs--
	fmt.Println(zs)
	//赋值运算
	zs +=3
	fmt.Println(zs)

	//浮点数,float32,float64,有精度损耗,
	var hight float64
	fmt.Println(hight)
	var hightt float64 = 3
	fmt.Printf("%T,%f\n",hightt,hight)

	fmt.Println(1.1+1.2)
	fmt.Println(1.1-1.2)
	fmt.Println(1.1*1.2)
	fmt.Println(1.1/1.2)
	hightt++
	fmt.Println(hightt)




	//字符
	//"" 可解释字符串，·· 原生字符串，特殊字符\n \r \f \v只能在"",''rune类型
	var _name string = "KK\tK"
	var _desc string = `胖\t子`
	fmt.Println(_name,_desc)

	//指针
	var A = 2
	BC := A
	BC = 4
fmt.Println(A,BC)
	var CC *int = &A
	C := &A
	fmt.Printf("%T,%T\n",CC,C)
	//用户输入
	var name string
	fmt.Println("请输入你的名字：",name)
	fmt.Scan(&name)
	fmt.Println("名字",name)

	var age int
	fmt.Println("请输入你的年龄：" , age)
	fmt.Scan(&age)
	fmt.Println("年龄:",age)
}