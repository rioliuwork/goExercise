package function

import (
	"fmt"
	"reflect"
)

func ReflectExercise()  {
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
	Sex int
	Age int `json:"age"`
}

func ReflectExercise2()  {
	typeOfMyStruct := reflect.TypeOf(myStruct{})
	fmt.Println(typeOfMyStruct.Name())
	fmt.Println(typeOfMyStruct.Kind())
}