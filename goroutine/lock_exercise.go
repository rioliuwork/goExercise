package goroutine

import (
	"fmt"
	"sync"
	"time"
)

//定义全局变量,表示粮食总量
var bread = 10

//创建同步等待组
var group sync.WaitGroup

//创建锁
var mutex sync.Mutex

func LockExercise() {
	group.Add(4)
	go Issue("猪八戒")
	go Issue("孙悟空")
	go Issue("沙和尚")
	go Issue("唐三藏")
	group.Wait()
}

//资源发放方法
func Issue(name string) {
	defer group.Done()
	for {
		//上锁
		mutex.Lock()
		if bread > 0 {
			//加锁控制之后,每次只允许一个协程进来,避免争抢
			bread--
			fmt.Println(name, "抢到粮食,还剩下", bread, "个")
		} else {
			//条件不满足也需要解锁,否则就会造成死锁其他协程不能执行
			mutex.Unlock()
			fmt.Println(name, "别抢了,没有粮食了...")
			break
		}
		mutex.Unlock()
	}
}

//创建读写锁,可以是他的指针类型
var rwMutex sync.RWMutex

var rwGroup sync.WaitGroup

//读写锁练习
func ReadWriteLockExercise() {
	rwGroup.Add(3)
	go read(1)
	go write(2)
	go read(3)
	rwGroup.Wait()
	fmt.Println("读写锁练习函数执行结束")
}

//读数据方法
func read(i int) {
	defer rwGroup.Done()
	fmt.Println("=====准备读数据=====")
	//读上锁
	rwMutex.RLock()
	fmt.Println("正在读取...", i)
	//读解锁
	rwMutex.RUnlock()
	fmt.Println("=====读取结束=====")
}

//写数据方法
func write(i int) {
	defer rwGroup.Done()
	fmt.Println("=====准备写数据=====")
	//写数据上锁
	rwMutex.Lock()
	fmt.Println("正在写数据...", i)
	time.Sleep(1 * time.Second)
	//写数据解锁
	rwMutex.Unlock()
	fmt.Println("=====写数据结束=====")
}
