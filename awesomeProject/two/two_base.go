package main

import (
	"fmt"
	"time"
)

func main()  {
	//for 循环
	 var _VRD = 1
	for _VRD <= 6 {
		_VRD++
		fmt.Println(_VRD)

	}

	for fl := 2 ; fl >= 6;fl=fl+3 {
		fmt.Println(fl)

	}

	for {
		fmt.Println("zoo")
		break
	}
	 for ss :=2 ; ss <=10 ; ss= ss + 8 {
		 fmt.Println("这里",ss)

	var desc = "我是小美女"
	for i,ch :=  range desc {
		fmt.Printf("%T %d %q \n",i,ch,ch,)
	}
	 }

	//if/else判断
	if 6/3 == 2 {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
	}
	var bo = 80
	if  bo <=60  {
		fmt.Println("及格")
	}else if  bo >=70 || bo >=100  {
		fmt.Println("棒")
	} else {
		fmt.Println("加油")
	}

	var yes string
	fmt.Println("有卖西瓜的吗？Y/N")
	fmt.Scan(&yes)
	if yes == "Y" || yes == "y" {
		fmt.Println("老婆的想法：十个包子，一个西瓜")
	} else {fmt.Println("老婆的想法：十个包子")}
	if yes == "Y"|| yes =="y" {
		fmt.Println("老公的想法：一个包子")
	} else {fmt.Println("老公的想法：十个包子")}

	var score = 8
	if score <=100 && score >=98 {
		fmt.Println("A")
	} else if score >= 70 && score <98 {
		fmt.Println("B")
	} else if score >=60 && score <70 {
		fmt.Println("C")
	} else {
		fmt.Println("加油哦！")
	}

	//swith/case
	i := 2
	fmt.Println("Write",i,"as")
	switch i {
	case 1 :
		fmt.Println("one")
	case 2 :
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Now().Weekday() :
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekend")
	}
	t := time.Now()
	switch  {
	case t.Hour() < 12 :
		fmt.Println("It's before noon ")
	default:
		fmt.Println("It's after noon")
	}

	var cla string
	fmt.Println("你看见西瓜了吗？Y/N")
	fmt.Scan(&cla)
	switch {
	case cla == "Y" || cla == "y" :
		fmt.Println("老婆的想法： 一个西瓜，十个包子")
	default:
		fmt.Println("老婆的想法：十个包子")
	}
	switch {
	case 	cla == "y" || cla == "Y" :
		fmt.Println("老公的想法：一个包子")
	default:
		fmt.Println("老公的想法： 十个包子")
	}

	var cj = 99
	switch  {
	case cj == 100 :
		fmt.Println("棒极了")
	case cj >80 && cj < 100 :
		fmt.Println("优秀")
	default:
		fmt.Println("加油哦")

	//goto 跳转

	}




}