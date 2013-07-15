package controllers

import (
	"fmt"
	"github.com/hermanschaaf/ironzebra/app/models"
	"github.com/hermanschaaf/ironzebra/app/routes"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"github.com/russross/blackfriday"
	"html/template"
	"labix.org/v2/mgo/bson"
)

type Blog struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Blog) List() revel.Result {
	collection := c.Database.C("posts")
	postList := []models.Post{}
	query := bson.M{"published": true}
	if c.Session["role"] == "admin" {
		query = nil
	}
	iter := collection.Find(query).Sort("-timestamp").Limit(5).Iter()
	iter.All(&postList)
	isAdmin := c.Session["role"] == "admin"
	return c.Render(postList, isAdmin)
}

func (c Blog) ListAll() revel.Result {
	collection := c.Database.C("posts")
	postList := []models.Post{}
	query := bson.M{"published": true}
	if c.Session["role"] == "admin" {
		query = nil
	}
	iter := collection.Find(query).Sort("-timestamp").Iter()
	iter.All(&postList)
	isAdmin := c.Session["role"] == "admin"
	return c.Render(postList, isAdmin)
}

func (c Blog) Show(id int, slugString string) revel.Result {
	// Collection Posts
	collection := c.Database.C("posts")
	isAdmin := c.Session["role"] == "admin"

	// Query the post by short id
	result := models.Post{}
	query := bson.M{"shortid": id}
	if isAdmin == false {
		query["published"] = true
	}
	err := collection.Find(query).One(&result)
	if err != nil {
		panic(err)
	}

	// if the slug is wrong, redirect to correct slug
	if result.Slug != slugString {
		fmt.Println(result.Slug)
		return c.Redirect(routes.Blog.Show(id, result.Slug))
	}
	if result.Slug == "" {
		return c.Redirect(routes.Blog.List())
	}

	// parse markdown into HTML
	html_output := template.HTML(string(blackfriday.MarkdownCommon([]byte(result.Body))))

	showAuthor := true
	rootUrl, _ := revel.Config.String("zebra.root_url")

	return c.Render(result, rootUrl, html_output, showAuthor, isAdmin)
}

func (c Blog) RedirectToSlug(id int) revel.Result {
	return c.Show(id, "")
}

func (c Blog) RedirectToPost(id int, slugString string) revel.Result {
	// redirect for users coming to the legacy /news url from google
	return c.Redirect(routes.Blog.Show(id, slugString))
}
