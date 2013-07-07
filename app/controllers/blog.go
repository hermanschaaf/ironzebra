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
	iter := collection.Find(nil).Iter()
	iter.All(&postList)
	return c.Render(postList)
}

func (c Blog) Show(id int, slug string) revel.Result {
	// Collection Posts
	collection := c.Database.C("posts")

	// Query the post by short id
	result := models.Post{}
	err := collection.Find(bson.M{"shortid": id}).One(&result)
	if err != nil {
		panic(err)
	}

	// if the slug is wrong, redirect to correct slug
	if result.Slug != slug {
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

	return c.Render(result, rootUrl, html_output, showAuthor)
}

func (c Blog) RedirectToSlug(id int) revel.Result {
	return c.Show(id, "")
}
