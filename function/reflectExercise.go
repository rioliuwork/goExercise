package function

import (
	"fmt"
	"reflect"
)

func ReflectExercise() {
	a := 36
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	aType := reflect.TypeOf(a)
	fmt.Println(aType.Name())
	aValue := reflect.ValueOf(a)
	fmt.Println(aValue.Int())
}

type myStruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func ReflectExercise2() {
	typeOfMyStruct := reflect.TypeOf(myStruct{})
	fmt.Println(typeOfMyStruct.Name())
	fmt.Println(typeOfMyStruct.Kind())
}

func ReflectValueExercise() {
	a := 2020
	valOf := reflect.ValueOf(a)
	fmt.Println(valOf.Interface())
	fmt.Println(valOf.Int())
}

func ReflectValueIsNilExercise() {
	//声明一个变量a为nil的空指针
	var a *int
	//判断是否为nil，返回true
	fmt.Println(reflect.ValueOf(a).IsNil())
	fmt.Println(reflect.ValueOf(a).IsValid())
	//当reflect.Value不包含任何信息,值为nil的IsValid()返回false
	fmt.Println(reflect.ValueOf(nil).IsValid())
}

func ReflectChangeValueExercise() {
	a := 100
	fmt.Printf("a的内存地址为:%p\n", &a)
	//获取变量a的反射类型reflect.Value的地址
	rf := reflect.ValueOf(&a)
	fmt.Println("通过反射获取变量a的地址:", rf)

	//获取a的地址的值
	val := rf.Elem()
	fmt.Println("反射a的值:", val)

	//修改a的值
	val.SetInt(200)
	fmt.Println("修改之后反射类型的值为:", val.Int())

	//原始值也被修改
	fmt.Println("原始a的值也被修改为:", a)
}

func add(a int, b int) int {
	return a + b
}

func ReflectFunctionExercise() {
	r_func := reflect.ValueOf(add)

	//设置函数需要传入的参数也必须是反射类型
	params := []reflect.Value{reflect.ValueOf(20), reflect.ValueOf(10)}

	//反射调用函数
	res := r_func.Call(params)

	//获取返回值
	fmt.Println(res[0].Int())
}
