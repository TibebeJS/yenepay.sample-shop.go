package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	Books := []struct {
		Title string
		Price float64
		Image string
	}{
		{
			"Go Bootcamp: 2nd Edition",
			14.99,
			"/public/img/go-bootcamp.jpg",
		},
		{
			"Automate The Boring Stuff With Python",
			24.99,
			"/public/img/automate-the-boring-stuff-with-python.jpg",
		},
		{
			"Get Programming With Go",
			99.99,
			"/public/img/get-programming-with-go.jpg",
		},
		{
			"Go Bootcamp: 2nd Edition",
			14.99,
			"/public/img/go-bootcamp.jpg",
		},
		{
			"Automate The Boring Stuff With Python",
			24.99,
			"/public/img/automate-the-boring-stuff-with-python.jpg",
		},
		{
			"Get Programming With Go",
			99.99,
			"/public/img/get-programming-with-go.jpg",
		},
	}

	return c.Render(Books)
}
