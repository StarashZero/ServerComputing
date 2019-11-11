package service

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func NewServer(port string) {
	m := martini.Classic()
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	}))
	m.RunOnAddr(":" + port)
}
