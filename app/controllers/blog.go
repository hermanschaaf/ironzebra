package controllers

import (
	"bitbucket.org/gosimple/slug"
	"fmt"
	"github.com/hermanschaaf/ironzebra/app/models"
	"github.com/hermanschaaf/ironzebra/app/routes"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"github.com/russross/blackfriday"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Blog struct {
	*revel.Controller
	revmgo.MongoController
}

func AddPost(collection mgo.Collection) {
	// Index
	index := mgo.Index{
		Key:        []string{"short_id", "title", "tags"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Dataz
	post_id := 100
	title := "Zebra, Gopher. Gopher, Zebra"
	subtitle := "A short and sweet zebra-gopher love story"
	body := "It was a normal day on the barren African plains deep inside of South Africa. It was dust and dirt speckled with some low shrubbery as far as the eye could see. It has been a dry "
	err = collection.Insert(&models.Post{ShortID: post_id, Title: title, Slug: slug.Make(title), Subtitle: subtitle, Body: body, Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}
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

	// parse markdown into HTML
	html_output := template.HTML(string(blackfriday.MarkdownCommon([]byte(result.Body))))

	showAuthor := true
	rootUrl, _ := revel.Config.String("zebra.root_url")

	return c.Render(result, rootUrl, html_output, showAuthor)
}

func (c Blog) RedirectToSlug(id int) revel.Result {
	return c.Show(id, "")
}
