// +build js,wasm

package main

import (
	"github.com/rajnikant12345/engine"
	"github.com/rajnikant12345/Resume/app"
)

func main() {
	c := make(chan struct{}, 1)

	println("WASM Go Initialized Now")

	app.RegisterCallbacks()
	engine.InitApp()
	engine.App.AddChild( app.CreateApp() )
	engine.StartApp()
	<-c
}
