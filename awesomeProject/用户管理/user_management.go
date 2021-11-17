package main

import (
	"fmt"
	"strconv"
	"strings"
)

//1添加用户
func add(pk int ,users map[string]map[string]string)  {
	var (
		id string = strconv.Itoa(pk)
		name string
		age string
		addr string
		tel string
	)

	fmt.Println("请输入你的名字")
	fmt.Scan(&name)
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	fmt.Println("请输入你的地址：")
	fmt.Scan(&addr)
	fmt.Println("请输入你的电话：")
	fmt.Scan(&tel)

	users[id] = map[string]string{
		"id": id ,
		"name": name,
		"age":age,
		"addr":addr,
		"tel":tel,
	}

}


//查询用户，
func  query (users  map[string]map[string]string) {
	var q string
	fmt.Println("请输入您要查询的数据:", q)
	fmt.Scan(&q)
	//定义标题，打印标题
	title := fmt.Sprintf("%5s|%10s|%5s|%10s|%50s", "id", "name", "age", "tel", "addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("_", len(title)))
	for _, user := range users {
		//strings.Contains遍历map中是否包含某个子字符串
		//查询用户信息，并显示
		if strings.Contains(user["name"], q) || strings.ContainsAny(user["tel"],q) || strings.Contains(user["addr"],q) {
			//%5s占5位，
			fmt.Printf("%5s|%10s|%5s|%10s|%50s",user["id"],user["name"],user["age"],user["tel"],user["addr"])
			fmt.Println()
		}
	}
}



func main()  {
	//存储用户信息
	users := make(map[string]map[string]string)
	id := 0
	//用户管理：用户信息存储，用户的增删改查
	fmt.Println("欢迎来到小芳元宇宙")
	for {
		var user1 string
		fmt.Println(`请输入以下操作：
						1 创建用户
						2 修改用户
						3 删除用户
						4 查询用户
						5 退出
`)


	fmt.Scan(&user1)
	fmt.Println("请输入你的指令：",user1)
	if user1 == "1" {
		id++
		add(id,users)

	} else if user1 == "2" {

	} else if user1 == "3" {

	} else if user1 == "4"{
		query(users)
	} else if user1 == "5" {
		break
	} else {
		fmt.Println("选择错误")
	}
}
}