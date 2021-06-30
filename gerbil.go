package gerbil

import (
	"strings"
	"syscall/js"
)

var App *Gerbil

type Gerbil struct {
	routes map[string]interface{}
}

func New() *Gerbil {
	return &Gerbil{
		routes: make(map[string]interface{}),
	}
}

func (*Gerbil) renderApplication(c interface{}) {
	dom := js.Global().Get("document")
	root := dom.Call("getElementById", "root")
	root.Set("innerHTML", c.(Component).Render())
}

func (*Gerbil) currentRoute() string {
	window := js.Global().Get("window")
	location := window.Get("location")
	rawUrl := location.Get("href").String()
	baseUrl := strings.Split(rawUrl, "://")[1]
	if len(strings.Split(baseUrl, "/?")) == 1 {
		return "/"
	}
	if strings.Split(baseUrl, "/?")[1] == "" {
		return "/"
	}
	return "/" + strings.Split(baseUrl, "/?")[1]
}

func (g *Gerbil) Run() {
	component, ok := g.routes[g.currentRoute()]
	if ok {
		g.renderApplication(component)
	} else {
		//TODO - render 404
	}

	<-make(chan struct{})
}

func (g *Gerbil) Route(path string, component interface{}) {
	if !IsComponent(component) {
		return
	}
	g.routes[path] = component
}
