package service

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func NewServer(port string) {
	//get a martini
	m := martini.Classic()
	//set assets path
	m.Use(martini.Static("assets"))
	//set render template
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	}))
	//deal with get request
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})
	//run martini
	m.RunOnAddr(":" + port)
}
