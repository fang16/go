package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

//数组
var a [6]int

func main()  {
	var w [4]int
	fmt.Printf("%\nT",w)

	var w1 [4]bool
	var w2 [4]string
	var w3 [4]byte
	fmt.Println(w1)
	fmt.Printf("%q\n",w2)
	fmt.Printf("%q\n",w3)

	//数组+字面量
	w = [4]int{10,20,30}
	fmt.Println(w)
	w = [4]int{1:10,2:40,3:50}
	fmt.Println(w)


	fmt.Println("emp:" ,a)
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	fmt.Println("len:",len(a))
	fmt.Println("len4:",len("a[4]"))

	bb := [5]int{1,2,3,4,5}
	fmt.Println(bb)

	var twoD [2][3]int
	for i :=0 ; i < 2 ; i++ {
		for j := 1; j < 3; j++ {
			twoD[i][j] = i + j
		}

		num := [10]int{10, 20, 30, 40}
		fmt.Println("num", num)
		num = [...]int{1, 2, 3, 4, ' ', 5, 0, 3, 2, 1}
		fmt.Println("num1", num)
		num3 := [5]int{22, 33, 44}
		fmt.Printf("%T,%#v\n", num3, num3)

		num4 := [15]int{2: 20, 4: 40, 9: 90, 12: 12}
		fmt.Printf("%T,%#v\n", num4, num4)

		fmt.Println("qiepian",num4[2:5])
		fmt.Println("zheli",num4[0:3:13])
		fmt.Printf("%T\n",num[2:5])

		//多维数组，第一个[]长度，第二个[]为类型，数组的类型必须是一致的
		 m2 := [2][3]int{}
		 fmt.Println(m2)
		 fmt.Printf("%T\n",m2)



		num5 := [4]int{1, 2, 3, 4}
		num6 := [4]int{4, 3, 2, 1}
		num7 := [4]int{1, 2, 4, 5}
		num8 := [4]int{1, 2, 3, 4}
		fmt.Println(num5 == num6)
		fmt.Println(num5 != num7)
		fmt.Println(num5 == num8)
		fmt.Println(len(num5))

		//遍历所有值
		for i := 0; i < len(num4);i++ {
			fmt.Println(num4[i])
		}
		for index,value := range num4 {
			fmt.Println(index,value)
		}

	}
	fmt.Println("two",twoD)

	//切片,数组的切片的出来的也是切片，[]int
	s := make( []string,3)
	fmt.Println("emp:",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set",s)
	fmt.Println("get",s[2])
	fmt.Println("len:",len(s))
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("append:",s)

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	d := s[2:5]
	fmt.Println(d)

	e := s[:5]
	fmt.Println(e)

	f := s[2:]
	fmt.Println("test:",f)

	twob := make([][]int,3)
	for i :=0 ; i<3; i++ {
		innerLen := i+1
		twob[i] =make([]int,innerLen)
		for j := 0 ; j<innerLen; j++ {
			twob[i][j]=i+j
		}
	fmt.Println("here",twob)
	}

	var q []int
	fmt.Printf("%T,%q\n",q,q)

	//字面量
	q =  []int{1,2,3}
	fmt.Println(q)

	q= []int{1,2,3,4}
	fmt.Println(q)

	//数组切片赋值

	var q1 [10]int = [10]int{1,2,3,4,5,6,4,0,3,2}
	q = q1[1:10]
	fmt.Println(q,len(q),cap(q))
	fmt.Printf("%T,%#v\n",q,q)

	//make函数,第一个参数代表长度，第二个代表容量，
	q = make([]int,3)
	fmt.Printf("%#v,%d,%d\n",q,q,q)

	q = make([]int,3,7)
	fmt.Printf("%#v,%d,%d\n",q,q,q)

	var q2 = []int{1,2,3,4,5}
	fmt.Println(q2)
	q2[2] = 10
	fmt.Println(q2)
	q2[4] = 40
	fmt.Println(q2)
	q2 = append(q2,2)
	fmt.Printf("%#v,%d,%T\n",q2,q2,q2)
	q2 = append(q,12,34,53,22,22,33,22,11,44)
	fmt.Println(q2)

	for _,value := range q2 {
		fmt.Println(value)
	}

	for i :=0; i < len(q2); i++ {
		fmt.Println(q2[i])
	}
	fmt.Println(q2[2:5],len(q2),cap(q2))

	q3 := q2[2:6:13]
	fmt.Println(q3,len(q3),cap(q3))

	q4 := []int{1,2,3,4}
	q5 := []int{6,54,4,3,2,1,9}

	copy(q5[:5],q5[3:])
	fmt.Println("here",q5)

	copy(q4,q5)
	fmt.Println(q5,q4)

	//数组是值类型

	slice01 := []int{1,2,3,4}
	slice02 := slice01
	slice02[2] = 99
	fmt.Println(slice02,slice01)

	array01 := [4]int{1,2,3,4}
	array02 := array01
	array02[2] = 99
	fmt.Println(array01,array02)

	//排序
	sort1 := []int{33,55,22,13,4,0}
	sort.Ints(sort1)
	fmt.Println(sort1)
	sort2 := []string{"ak","ka","test","sat"}
	sort.Strings(sort2)
	fmt.Println(sort2)


	//映射，定义完映射之后必须把值初始化之后才能操作,无序的

	var nil1 map[string]int //映射
	fmt.Printf("%T,%#v \n",nil1,nil1)
	fmt.Println(nil1 == nil)

	nil1 = map[string]int{"小芳":100,"张林源":99,"老牛":0}
	fmt.Printf("%T,%#v,%d\n",nil1,nil1,nil1)
	fmt.Println(nil1)

	//增删改查
	fmt.Println(nil1["小芳"])
	fmt.Println(nil1["陈林"])
	if v,ok := nil1["陈林"]; ok {
		fmt.Println(v)
	}
	if v , ok := nil1["小芳"]; ok {
		fmt.Println(v)
	}

	nil1["陈林"] = -1
	fmt.Println(nil1)

	nil1["陈林"] = 002
	fmt.Println(nil1)

	delete(nil1,"陈林")
	fmt.Println(nil1)
	fmt.Println(len(nil1))

	for k,v := range nil1 {
		fmt.Println(k,v)
	}

	nil2 := nil1
	fmt.Println(nil2)
	nil2["陈林"]=101
	fmt.Println(nil2)
	fmt.Println(nil1)

	nil3 := map[string]string{"南昌":"上饶","深圳":"宝安","北京":"朝阳"}
	fmt.Println(nil3)

	var nil4 map[string]map[string]string
	nil4 = map[string]map[string]string{"小芳":{"地址":"南昌","电话":"123432","现住址":"深圳"},"张林源":{"地址":"山西","电话":"1234567"}}
	fmt.Printf("%T,%#v\n",nil4,nil4)

	nil4["张林源"]["地址"] = "江西"
	fmt.Println(nil4)

	//函数
	fmt.Println(strings.Compare("abc","bac"))
	fmt.Println(strings.Contains("abc","bef"))
	fmt.Println(strings.ContainsAny("abc","befd"))
	fmt.Printf("%q\n",strings.Fields("abc def\ncde\rdefkjl\fkiolsdj\vdejindsk"))

	//字节切片
	cc := []byte{'q','b','c','d'}
	fmt.Printf("%T,%v\n",cc,cc)
	ss := string(cc)
	fmt.Printf("%T,%v\n",ss,ss)

	da := "我爱张林源"
	fmt.Println(len(da))
	fmt.Println(utf8.RuneCountInString(da))

	//字符串转换
	//字符串到其他类型转换
	//==》boot
	if v , err := strconv.ParseBool("true"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	//==>> int
	if v , err := strconv.Atoi("12765432"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	if v , err := strconv.ParseInt("23",18,15); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	//float
	if v , err := strconv.ParseFloat("1.4",43); err == nil {
		fmt.Printf("%T,%#v\n",v,v)
	} else {
		fmt.Println(err)
	}

	//






}
