package controllers

import (
	"github.com/hermanschaaf/ironzebra/app/models"
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
