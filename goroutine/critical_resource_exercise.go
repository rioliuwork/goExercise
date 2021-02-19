package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//全局变量 表示食物总数
var food = 10

func CriticalResource() {
	//开启4个线程抢资源
	go Relief("孙悟空")
	go Relief("猪八戒")
	go Relief("沙和尚")
	go Relief("唐三藏")

	//设置休眠时间,让所有子协程都执行完
	time.Sleep(3 * time.Second)
}

//定义一个食物发放方法
func Relief(name string) {
	for {
		if food > 0 { //此时有可能第二个goroutine访问的时候,第一个goroutine还未执行完,所以条件也成立
			//设置随机休眠时间
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			food--
			fmt.Println(name, "抢到粮食,还剩下", food, "个")
		} else {
			fmt.Println(name, "别抢了，没有粮食了...")
			break
		}
	}
}

//创建一个同步等待组的对象
var wg sync.WaitGroup

func WaitGroupExercise() {
	//设置同步等待组的数量
	wg.Add(3)
	go func1()
	go func2()
	go func3()
	//当前goroutine进入阻塞状态
	wg.Wait()
	fmt.Println("main end...")
}

func func1() {
	fmt.Println("func1 done...")
	//执行完成,同步等待锁数量减1
	wg.Done()
}

func func2() {
	defer wg.Done()
	fmt.Println("func2 done...")
}

func func3() {
	//推荐使用延时执行的方法来减去执行组的数量
	defer wg.Done()
	fmt.Println("func3 done...")
}
