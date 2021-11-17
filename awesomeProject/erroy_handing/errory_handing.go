package main

import (
	"errors"
	"fmt"
)
//错误处理
//返回值：error定义错误类型
//errors.Now(),fmt.Errorf() 创建错误返回值
//无错误 nil

func division( a ,b int) (int ,error)  {
	if b == 0 {
		return -1,errors.New("division by zore")
	}

	return a/b,nil
}
//常用于处理第三方库
func test() (err error)  {
	if d := recover(); d != nil {
		fmt.Println("error test")
		err = fmt.Errorf("%v\n",d)
	}
	panic("error")
	return
}


func main()  {
	fmt.Println(division(1,3))
	fmt.Println(division(1,0))
	if v,err := division(2,0);err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	e := fmt.Errorf("error: %s","by division zore")
	fmt.Printf("%T,%v\n",e,e)

	//延时执行,谁先注册谁后执行,堆栈的应用
	defer func() {
		fmt.Println("defer01")
	}()
	defer func() {
		fmt.Println("defer02")
	}()

	fmt.Println("main defer")

	//panic
	fmt.Println("main start")
	panic("error")
	fmt.Println("over")

	defer func() {
		if c := recover(); c != nil {
			fmt.Println(c)
		}
	}()

	fmt.Println(recover())

	//
	err := test()
	fmt.Println(err)
}