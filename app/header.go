package app

import (
	"github.com/rajnikant12345/engine"
	"syscall/js"
)

func CreateNav(name string) *engine.Element  {
	a := engine.NewElement("a")
	a.SetClass("f6 f5-l link bg-animate black-80 hover-bg-lightest-blue dib pa3 ph4-l")
	a.SetInnerHtml(name)

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

func CreateHeader() *engine.Element {

	e := engine.NewElement("div")
	h1 := engine.NewElement("header")
	h1.SetClass("tc pv1")
	h1.SetInnerHtml(`<header class="tc pv1">
  <img src="./rajni.JPG" class="br-100 pa1 ba b--black-10 h4 w4" alt="avatar">
  <h1 class="f5 f4-ns fw6 mid-gray">Rajni Kant</h1>
  <h2 class="f6 gray fw2 ttu tracked">Software Devloper, Noida, India</h2>
  <h2 class="f6 gray fw2">Email: rajnikant12345@live.com </h2>
  <a class="f5 fw7 dib no-underline bg-animate bg-white hover-bg-light-blue dark-gray" href="https://www.linkedin.com/in/rajni-kant-94b90a18">My LinkedIn Profile</a>
</header>`)
	e.AddChild(h1)

	nav := engine.NewElement("nav")
	nav.SetClass("bt bb tc mw7 center mt4")
	nav.AddChild(CreateNav("Home"))
	nav.AddChild(CreateNav("Skills"))
	nav.AddChild(CreateNav("Organizations-Worked"))
	//nav.AddChild(CreateNav("Contact"))

	e.AddChild(nav)
	return e
}

