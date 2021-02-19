package goroutine

import "fmt"

func ChannelExercise() {
	var channel chan int
	fmt.Printf("通道数据类型:%T,通道的值:%v\n", channel, channel)

	if channel == nil {
		channel = make(chan int)
		fmt.Printf("通过make创建的通道数据类型:%T,通道的值:%v\n", channel, channel)
	}
}

func ChannelUseExercise() {
	channel := make(chan int)
	go func() {
		fmt.Println("=====子协程执行=====")
		data := <-channel
		fmt.Println("读取到通道中的数据是:", data)
	}()
	//往通道里放数据
	channel <- 10
	fmt.Println("=====主协程结束=====")
}

func ChannelCloseExercise1() {
	channel := make(chan int)
	go func() {
		fmt.Println("=====子协程执行=====")
		for i := 0; i < 10; i++ {
			channel <- i
		}
		//关闭通道
		close(channel)
	}()

	for {
		//获取通道的状态和数据
		v, ok := <-channel
		if !ok {
			fmt.Println("子协程已将通道关闭")
			break
		}
		fmt.Println("获取到子协程数据为", v)
	}

	fmt.Println("主协程结束")
}

func ChannelCloseExercise2() {
	channel := make(chan int)

	go func() {
		fmt.Println("=====子协程执行=====")
		for i := 0; i < 10; i++ {
			channel <- i
		}
		//关闭通道
		close(channel)
	}()

	//通过for range循环读取通道中的数据,当通道关闭,循环也就结束了
	for v := range channel {
		fmt.Println("获取到子协程发送的数据为", v)
	}

	fmt.Println("主协程结束")
}

func BufferChannelExercise() {
	//定义一个缓冲区大小为5的通道
	channel := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("子协程写入数据:", i)
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println("主协程读取到的数据为:", v)
	}

	fmt.Println("主协程执行结束")
}

//读写通道练习
func ReadWriteChannelExercise() {
	//双向通道
	channel := make(chan int)
	writeOnlyChannel := make(chan<- int)
	readOnlyChannel := make(<-chan int)

	//如果创建的双向通道
	//则在子协程内部写入数据,主协程读取的时候不受影响
	go writeOnly(channel)
	go ReadeOnly(channel)

	//如果创建的只读通道,则不能读取写回来的数据
	go writeOnly(writeOnlyChannel)

	//创建的只读通道,无法向通道里面写数据
	go ReadeOnly(readOnlyChannel)

	fmt.Println("结束")
}

//只读
func ReadeOnly(readeOnlyChannel <-chan int) {
	data := <-readeOnlyChannel
	fmt.Println("读取到的通道数据是:", data)
}

//只写
func writeOnly(writeOnlyChannel chan<- int) {
	//如果传进来的原本是双向通道
	//但函数接收的是一个只写通道,则在此函数中只允许写数据,不允许读数据
	//单向通道往往是作为参数传递
	writeOnlyChannel <- 10
	fmt.Println("只写通道结束")
}
