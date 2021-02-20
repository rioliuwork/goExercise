package file

import (
	"fmt"
	"os"
)

func FileExercise() {
	//通过os.Stat获取文件的信息
	fileinfo, err := os.Stat("D:\\GoExercise\\leetCode\\materials\\testFile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileinfo.Name())
}

//创建文件夹
func CreateDirExercise() {
	//Mkdir创建文件夹
	err := os.Mkdir("testdir", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	//MkdirALl创建多级文件夹
	err2 := os.MkdirAll("testdir/test/aa", os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
}

//创建文件
func CreateFileExercise() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
}

func ReadFileExercise() {
	//通过Open只读方式打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Name())
	file.Close()

	/**
	const (
		O_RDONLY int = syscall.O_RDONLY // 以只读方式打开文件
		O_WRONLY int = syscall.O_WRONLY //  以只写方式打开文件
		O_RDWR   int = syscall.O_RDWR   //  以读写方式打开文件
		O_APPEND int = syscall.O_APPEND // 写入时向文件追加数据
		O_CREATE int = syscall.O_CREAT  // 如果不存在创建新文件
		O_EXCL   int = syscall.O_EXCL //与O_CREATE一起使用，文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 为同步I/O打开
		O_TRUNC  int = syscall.O_TRUNC  // 打开时截断常规可写文件
	)
	*/
	file2, err2 := os.OpenFile("test2.txt", os.O_RDONLY|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file2.Close()
	fmt.Println(file2)
}

func RemoveFile() {
	err := os.Remove("test2.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveAllFile() {
	err := os.RemoveAll("testdir")
	if err != nil {
		fmt.Println(err)
	}
}
