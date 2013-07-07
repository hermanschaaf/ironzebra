package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/hermanschaaf/ironzebra/app/models"
	"github.com/hermanschaaf/ironzebra/app/routes"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
)

type Admin struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) getUser(username string) *models.User {
	users := c.Database.C("users")
	result := models.User{}
	users.Find(bson.M{"username": username}).One(&result)
	return &result
}

func (c Admin) Login(username, password string) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Admin.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Admin.Index())
}
