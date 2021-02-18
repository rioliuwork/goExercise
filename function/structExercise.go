package function

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string  `json:"NAME"`
	Age  int     `json:"AGE"`
	Son  *Person `json:"SON,omitempty"`
}

func GetPersonJsonInfo() {
	p := Person{}
	p.Name = "Go大宝"
	p.Age = 30
	p.Son = &Person{
		Name: "Go小宝",
		Age:  5,
	}
	buf, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func JsonToStruct() {
	jsonStr := `{"name":"Go大宝","age":30,"son":{"name":"Go小宝","age":5}}`
	var p Person
	if err := json.Unmarshal([]byte(jsonStr), &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
