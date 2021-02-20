package file

import (
	"fmt"
	"io"
	"os"
)

func ReadeData() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]byte, 30, 30)
	n, e := file.Read(data)
	if e != nil {
		fmt.Println(e)
	}
	//读取到的字节数量
	fmt.Println(n)
	//读取到的数据内容转换为string
	fmt.Println(string(data))
}

func WriteData() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "这是写入的数据"
	b := []byte(str)

	//写到文件
	n, e := file.Write(b)
	if e != nil {
		fmt.Println(e)
	}
	//写入到文件的字节数
	fmt.Println(n)
}

func CopyFile() {
	//源文件
	file1, err1 := os.Open("test.txt")
	if err1 != nil {
		fmt.Println(err1)
	}

	//目标文件
	file2, err2 := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file1.Close()
	defer file2.Close()

	byteStr := make([]byte, 1024, 1024)
	for {
		n, e := file1.Read(byteStr)
		if e == io.EOF || n == 0 {
			fmt.Println("读取文件结束")
			break
		}
		//循环写入文件
		file2.Write(byteStr[:n])
	}
}

func CopyFileByOs() {
	//源文件
	file1, err1 := os.Open("test.txt")
	if err1 != nil {
		fmt.Println(err1)
	}

	//目标文件
	file2, err2 := os.OpenFile("test3.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file1.Close()
	defer file2.Close()

	n, e := io.Copy(file2, file1)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(n)
}
