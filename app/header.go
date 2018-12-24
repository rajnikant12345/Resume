package app

import (
	"github.com/rajnikant12345/engine"
	"syscall/js"
)

func CreateNav(name string) *engine.Element  {
	a := engine.NewElement("a").SetClass("f6 f5-l link bg-animate black-80 hover-bg-lightest-blue dib pa3 ph4-l").SetInnerHtml(name)

	cb := js.NewCallback(func(args []js.Value) {

		Body.RemoveChild()
		if name == "Home" {
			Body.AddChild(Home)
		}
		if name == "Skills" {
			Body.AddChild(Skills)
		}
		if name == "Organizations-Worked" {
			Body.AddChild(Orgs)
		}
		if name == "Contact" {
			Body.AddChild(Skills)
		}


	})
	a.SetCallBack("click",cb)
	return a
}

func CreateImage(src , class string) *engine.Element {
	return engine.NewElement("img").Set("src",src).SetClass(class)
}

func CreateHeading(typ , value, class string ) *engine.Element {
	return engine.NewElement(typ).SetInnerHtml(value).SetClass(class)
}

func CreateAnchor( class, href , text string) *engine.Element {
	return engine.NewElement("a").SetClass(class).Set("href",href).SetInnerHtml(text)
}

func CreateHeader() *engine.Element {

	e := engine.NewElement("div")
	h1 := engine.NewElement("header").SetClass("tc pv0").Nest(
		CreateImage("./rajni.JPG","br-100 pa1 ba b--black-10 h4 w4"),
		CreateHeading("h1","Rajni Kant", "f5 f4-ns fw6 mid-gray"),
		CreateHeading("h2","Software Devloper, Noida, India", "f6 gray fw2 ttu tracked"),
		CreateHeading("h2","Email: rajnikant12345@live.com", "f6 gray fw2"),
		CreateAnchor("f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray",
			"https://www.linkedin.com/in/rajni-kant-94b90a18", "My LinkedIn Profile").Set("target","_blank"),
	)

	e.AddChild(h1)

	nav := engine.NewElement("nav").SetClass("bt bb tc mw7 center mt4").Nest(
		CreateNav("Home"),
		CreateNav("Skills"),
		CreateNav("Organizations-Worked"))

	e.AddChild(nav)
	return e
}

