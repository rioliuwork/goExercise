package function

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string
	Sex  int
	Age  int
}

func ReflectPoint() {
	typeOfMyStruct := reflect.TypeOf(&MyStruct{})
	fmt.Println(typeOfMyStruct.Elem().Name())
	fmt.Println(typeOfMyStruct.Elem().Kind())
}
