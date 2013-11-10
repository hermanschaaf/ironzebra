package controllers

import (
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
)

func (c Admin) checkUser() revel.Result {
	_, loggedIn := c.Session["username"]
	if loggedIn == false {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	} else {
		isAdmin := c.Session["role"] == "admin"
		if isAdmin == false {
			c.Flash.Error("Please log in as an admin first")
			return c.Redirect(App.Login)
		}
	}
	return nil
}

func init() {
	revel.ERROR_CLASS = "error"
	revmgo.ControllerInit()
	revel.InterceptMethod(Admin.checkUser, revel.BEFORE)
}
