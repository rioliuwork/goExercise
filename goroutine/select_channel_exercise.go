package goroutine

import (
	"fmt"
	"time"
)

func SelectChannelExercise() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go func() {
		for {
			select {
			case <-channel1:
				fmt.Println("成功获取到channel1的数据:", <-channel1)
			case channel2 <- 1:
				fmt.Println("成功向channel2中写入数据")
			case <-time.After(time.Second * 2):
				//使用time.After 设置超时响应,如果迟迟接收不到以上的case,就会响应超时
				fmt.Println("超时")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		channel1 <- i
		fmt.Println("channel1写入数据:", i)
	}

	for i := 0; i < 10; i++ {
		info := <-channel2
		fmt.Println("获取到channel2的数据:", info)
	}

	//select 会一直等待某个 case 语句完成
	//成功从 channel1 或者 channel2 中读到数据,则select语句结束
}
