package controllers

import (
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

func getPosts(c Blog, limit int, category string, admin bool) []models.Post {
	collection := c.Database.C("posts")
	postList := []models.Post{}
	query := bson.M{"published": true}
	if c.Session["role"] == "admin" {
		query = bson.M{}
	}
	if category != "" {
		query["category"] = category
	}
	q := collection.Find(query).Sort("-timestamp")
	if limit > 0 {
		iter := q.Limit(limit).Iter()
		iter.All(&postList)
		return postList
	} else {
		iter := q.Iter()
		iter.All(&postList)
		return postList
	}
}

func getCategories(c Blog) []models.Category {
	categoryList := []models.Category{}
	collection := c.Database.C("categories")
	iter := collection.Find(nil).Sort("name").Iter()
	iter.All(&categoryList)
	return categoryList
}

func (c Blog) List() revel.Result {
	isAdmin := c.Session["role"] == "admin"
	postList := getPosts(c, 5, "", isAdmin)
	categoryList := getCategories(c)
	return c.Render(postList, categoryList, isAdmin)
}

func (c Blog) ListAll() revel.Result {
	isAdmin := c.Session["role"] == "admin"
	postList := getPosts(c, 0, "", isAdmin)
	categoryList := getCategories(c)
	return c.Render(postList, categoryList, isAdmin)
}

func (c Blog) ListCategory(category string) revel.Result {
	isAdmin := c.Session["role"] == "admin"
	postList := getPosts(c, 0, category, isAdmin)
	categoryList := getCategories(c)
	currentCategory := category
	return c.Render(postList, categoryList, isAdmin, currentCategory)
}

func (c Blog) Show(category string, id int, slugString string) revel.Result {
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

	// if wrong category, redirect to correct category
	if result.Category != category {
		cat := result.Category
		if result.Category == "" {
			cat = "cats"
		}
		return c.Redirect(routes.Blog.Show(cat, id, result.Slug))
	}
	// if the slug is wrong, redirect to correct slug
	if result.Slug != slugString {
		return c.Redirect(routes.Blog.Show(category, id, result.Slug))
	}
	// TODO: a bit of a hack again..
	if result.Slug == "" {
		return c.Redirect(routes.Blog.List())
	}

	// parse markdown into HTML
	html_output := template.HTML(string(blackfriday.MarkdownCommon([]byte(result.Body))))

	showAuthor := true
	rootUrl, _ := revel.Config.String("zebra.root_url")

	canonical := routes.Blog.Show(category, id, result.Slug)
	return c.Render(result, canonical, rootUrl, html_output, showAuthor, isAdmin)
}

func (c Blog) RedirectToSlug(category string, id int) revel.Result {
	return c.Show(category, id, "")
}
