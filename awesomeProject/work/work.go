package main

import (
	"fmt"
	"time"
)
import "math/rand"

func main()  {
	for i := 0 ; i <= 9; i++ {

		for j :=1; j <= i;j++ {
			fmt.Printf("%d*%d=%d\t",j,i,j*i)
		}
		fmt.Println()
	}

flag:
	for d := 1 ; d < 6 ; d++ {
		//设置随机数种子
		rand.Seed(time.Now().Unix())
		//生成随机数
		a := rand.Intn(100)
		var work int
		fmt.Println("请猜猜输出的是什么数：")
		fmt.Scan(&work)
		fmt.Println(a)

		if d == 5 && a != work {
			fmt.Println("太笨了！请重新开始")
			goto flag

		}

		if a == work {
			fmt.Println("very goodA!!")
			break
		} else if a < work {
			fmt.Printf("太大了，请重猜,你还有%d次机会数量 \n",5-d)

		} else if a > work {
			fmt.Printf("太小了，请重猜,你还有%d次机会数量 \n", 5-d)

		}
	}



}

