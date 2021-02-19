package function

import (
	"fmt"
	"reflect"
)

type PersonInfo struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func ReflectFieldExercise() {
	personReflect := reflect.TypeOf(PersonInfo{})
	//获取结构体成员数量
	numField := personReflect.NumField()
	for i := 0; i < numField; i++ {
		fmt.Println("==================")
		//索引对应的字段信息
		field := personReflect.Field(i)
		fmt.Println(field)
		//返回对应的字段信息
		name, err := personReflect.FieldByName("Name")
		fmt.Println(name, err)
	}
}

func ReflectFieldValueExercise() {
	goValue := reflect.ValueOf(PersonInfo{
		Name: "Go小山",
		Sex:  0,
		Age:  30,
	})
	for i := 0; i < goValue.NumField(); i++ {
		fmt.Println("=====================")
		fmt.Println(goValue.Field(i).Interface())
		fmt.Println(goValue.Field(i).Type())
	}
	fmt.Println("=====================")

}
