package function

import "fmt"

type Role struct {
	Name    string
	Ability string
	Level   int
	Kill    float64
}

func (role Role) KongFu() {
	fmt.Printf("我是:%s,我的武功：%s,已经练到%d级了,杀伤力%.1f\n", role.Name, role.Ability, role.Level, role.Kill)
}
