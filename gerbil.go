package gerbil

import (
	"fmt"
	"reflect"
	"syscall/js"
)

var App *Gerbil

type Gerbil struct {
}

func New() *Gerbil {
	return &Gerbil{}
}

func (g *Gerbil) RenderComponent(c interface{}) {

	iface := reflect.TypeOf((*Component)(nil)).Elem()
	component := reflect.ValueOf(c)
	if component.Kind() != reflect.Ptr {
		fmt.Printf("'%s' must be a pointer to a struct that implements the '*gerbil.Component' interface\n", component.Type())
		return
	}

	if !component.Type().Implements(iface) {
		fmt.Printf("error: component %v does not implement %v\n", component.Type(), iface)
	}

	dom := js.Global().Get("document")
	root := dom.Call("getElementById", "root")
	root.Set("innerHTML", c.(Component).Render())
}
