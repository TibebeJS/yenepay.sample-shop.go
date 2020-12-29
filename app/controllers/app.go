package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	Books := []struct {
		Title      string
		Desciption string
		Image      string
	}{
		{
			"Go",
			"Learn Go",
			"/public/img/go-bootcamp.jpg",
		},
		{
			"Automate The Boring Stuff With Python",
			"Learn Go",
			"/public/img/automate-the-boring-stuff-with-python.jpg",
		},
		{
			"Get Programming With Go",
			"Learn Go",
			"/public/img/get-programming-with-go.jpg",
		},
		{
			"Go",
			"Learn Go",
			"/public/img/go-bootcamp.jpg",
		},
		{
			"Automate The Boring Stuff With Python",
			"Learn Go",
			"/public/img/automate-the-boring-stuff-with-python.jpg",
		},
		{
			"Get Programming With Go",
			"Learn Go",
			"/public/img/get-programming-with-go.jpg",
		},
	}

	return c.Render(Books)
}
