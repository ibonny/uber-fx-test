package module_one

import "fmt"

type ModuleOne struct {
}

func NewModuleOne() *ModuleOne {
	return &ModuleOne{}
}

func (mo *ModuleOne) SomeFunc() {
	fmt.Println("This is a test.")
}
