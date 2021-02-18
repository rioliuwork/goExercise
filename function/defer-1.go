package function

import "fmt"

func DeferExercise()  {
	defer func() {
		ms := recover()
		fmt.Println(ms, "恢复执行。。。")
	}()
	defer fmt.Println("第1个被defer执行")
	defer fmt.Println("第2个被defer执行")

	for i := 0; i <= 6; i++ {
		if i == 4 {
			panic("中断操作")
		}
	}

	defer fmt.Println("第3个被defer执行")
}
