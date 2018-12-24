package app

import "github.com/rajnikant12345/engine"

var Home *engine.Element
var Contact *engine.Element
var Orgs *engine.Element
var Skills *engine.Element
var Body *engine.Element

func CreateScroll( )  *engine.Element  {
	e := engine.NewElement("div")
	e.Node.Get("style").Set("overflowY","scroll")
	e.Node.Get("style").Set("height","400px")

	return e
}

func CreateApp() *engine.Element{

	Home = CreateHome()
	Contact = CreateContact()
	Orgs = CreateOrgs()
	Skills = CreateSkills()

	e := engine.NewElement("div")
	hdr := CreateHeader()

	Body = CreateScroll()
	Body.AddChild(Home)
	e.Nest(hdr ,Body)

	return e

}
