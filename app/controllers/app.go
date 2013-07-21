package controllers

import (
	"github.com/hermanschaaf/ironzebra/app/models"
	"github.com/hermanschaaf/ironzebra/app/routes"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

func (c App) Index() revel.Result {
	collection := c.Database.C("posts")
	result := models.Post{}
	collection.Find(bson.M{"published": true}).Sort("-timestamp").One(&result)
	return c.Render(result)
}

func (c App) Login() revel.Result {
	/*
		We define the admin login page here, because it's the
		only page on the admin interface that doesn't require
		the user to be logged in. This could probably be improved.
	*/
	_, loggedIn := c.Session["username"]
	if loggedIn {
		return c.Redirect(routes.Admin.Index())
	}
	return c.Render()
}
