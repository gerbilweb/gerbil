package gerbil

import (
	"fmt"
	"reflect"
)

type Component interface {
	Render() string
}

func IsComponent(c interface{}) bool {

	iface := reflect.TypeOf((*Component)(nil)).Elem()
	component := reflect.ValueOf(c)
	if component.Kind() != reflect.Ptr {
		fmt.Printf("'%s' must be a pointer to a struct that implements the '*gerbil.Component' interface\n", component.Type())
		return false
	}

	if !component.Type().Implements(iface) {
		fmt.Printf("error: component %v does not implement %v\n", component.Type(), iface)
		return false
	}

	return true
}
