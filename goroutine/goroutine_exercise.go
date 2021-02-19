package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func GoroutineExercise_1() {
	//使用关键字go调用函数或者方法,开启一个goroutine
	go testGo()
	for i := 0; i < 10; i++ {
		fmt.Println("主函数循环执行次数", i)
	}
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("函数执行结束")
}

func testGo() {
	for i := 0; i < 10; i++ {
		fmt.Println("测试goroutine", i)
	}
}

//匿名函数创建Goroutine
func AnonymousGoroutine() {
	go func() {
		fmt.Println("匿名函数创建goroutine执行")
	}()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("主函数执行")
}

//设置并发cpu核心数;中止当前goroutine联系代码
func SetGoMaxProcs() {
	//获取当前GOROOT目录
	fmt.Println("GOROOT:", runtime.GOROOT())
	//获取当前操作系统
	fmt.Println("操作系统:", runtime.GOOS)
	//获取当前逻辑CPU数量
	fmt.Println("逻辑CPU数量:", runtime.NumCPU())

	//设置最大的可同时使用的CPU核数 取逻辑cpu数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	//一般在使用之前就将cpu数量设置好（所以最好放在init函数内执行）
	fmt.Println(n)

	//goexit 终止当前goroutine
	//创建一个goroutine
	go func() {
		fmt.Println("goroutine start...")
		//runtime.Goexit()不能放在主goroutine(也就是主函数中)执行,否者会发生运行时恐慌
		runtime.Goexit()
		fmt.Println("goroutine end...")
	}()
	//休眠3秒,让子goroutine执行完
	time.Sleep(2 * time.Second)
	fmt.Println("function end...")
}
