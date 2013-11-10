package controllers

import (
	"bitbucket.org/gosimple/slug"
	"github.com/hermanschaaf/ironzebra/app/models"
	"github.com/hermanschaaf/ironzebra/app/routes"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Admin struct {
	*revel.Controller
	revmgo.MongoController
}

func addPost(database *mgo.Database, collection *mgo.Collection, title, subtitle, slug, category, body, image string) (post models.Post) {
	// Index
	index := mgo.Index{
		Key:        []string{"shortid", "timestamp", "title"},
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
	err = collection.Insert(&models.Post{
		ShortID:   models.GetNextSequence(database),
		Title:     title,
		Category:  category,
		Slug:      slug,
		Subtitle:  subtitle,
		Body:      body,
		Timestamp: time.Now(),
		Published: false})

	if err != nil {
		panic(err)
	}

	result := models.Post{}
	collection.Find(bson.M{"title": title}).One(&result)
	return result
}

func savePost(collection *mgo.Collection, shortID int, title, subtitle, slugString, category, body, image string) (post models.Post) {
	// Update Dataz
	err := collection.Update(bson.M{"shortid": shortID}, bson.M{
		"$set": bson.M{
			"title":    title,
			"subtitle": subtitle,
			"category": category,
			"slug":     slugString,
			"body":     body,
			"image":    image,
		},
	})

	if err != nil {
		panic(err)
	}

	result := models.Post{}
	collection.Find(bson.M{"title": title}).One(&result)
	return result
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Admin.Index())
}

func (c Admin) Edit(id int, slug string) revel.Result {
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
		return c.Redirect(routes.Blog.Show(result.Category, id, result.Slug))
	}

	// TODO: this is a bit of a hack
	if result.Slug == "" {
		return c.Redirect(routes.Admin.Index())
	}

	newPost := !result.Published

	categoryList := []models.Category{}
	collection = c.Database.C("categories")
	iter := collection.Find(nil).Sort("name").Iter()
	iter.All(&categoryList)

	return c.Render(result, categoryList, newPost)
}

func (c Admin) New() revel.Result {
	/* Create a new post */
	categoryList := []models.Category{}
	collection := c.Database.C("categories")
	iter := collection.Find(nil).Sort("name").Iter()
	iter.All(&categoryList)
	newPost := true
	return c.Render(categoryList, newPost)
}

func validatePost(c Admin, title, body, slugString, category string) {
	c.Validation.Required(title).Message("A title is required")
	c.Validation.Required(body).Message("You probably want some text in your post, no?")
	c.Validation.Required(slugString).Message("You need a slug...")
	c.Validation.Required(category).Message("You need to choose a category")
}

func (c Admin) SaveNew(title, subtitle, category, body, image string) revel.Result {
	slugString := slug.Make(title)
	validatePost(c, title, body, slugString, category)
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.New())
	}
	collection := c.Database.C("posts")
	result := addPost(c.Database, collection, title, subtitle, slugString, category, body, image)
	return c.Redirect(routes.Blog.Show(result.Category, result.ShortID, result.Slug))
}

func (c Admin) Save(id int, title, subtitle, slugString, category, body, publish, image string) revel.Result {
	validatePost(c, title, body, slugString, category)
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.Edit(id, slugString))
	}
	collection := c.Database.C("posts")
	if slugString == "" {
		slugString = slug.Make(title)
	}
	result := savePost(collection, id, title, subtitle, slugString, category, body, image)
	return c.Redirect(routes.Blog.Show(result.Category, result.ShortID, result.Slug))
}

func (c Admin) RedirectToSlug(id int) revel.Result {
	return c.Edit(id, "")
}

func (c Admin) Publish(id int) revel.Result {
	/* Publish the post */
	collection := c.Database.C("posts")
	collection.Update(bson.M{"shortid": id, "published": false}, bson.M{"$set": bson.M{"published": true, "timestamp": time.Now()}})

	result := models.Post{}
	collection.Find(bson.M{"shortid": id}).One(&result)
	return c.Redirect(routes.Blog.Show(result.Category, result.ShortID, result.Slug))
}

func (c Admin) Unpublish(id int) revel.Result {
	/* Unpublish the post */
	collection := c.Database.C("posts")
	collection.Update(bson.M{"shortid": id}, bson.M{"$set": bson.M{"published": false}})

	result := models.Post{}
	collection.Find(bson.M{"shortid": id}).One(&result)
	return c.Redirect(routes.Blog.Show(result.Category, result.ShortID, result.Slug))
}

func (c Admin) Categories() revel.Result {
	categoryList := []models.Category{}
	collection := c.Database.C("categories")
	iter := collection.Find(nil).Sort("name").Iter()
	iter.All(&categoryList)
	return c.Render(categoryList)
}

func (c Admin) NewCategory(name, description string) revel.Result {
	collection := c.Database.C("categories")
	err := collection.Insert(&models.Category{
		Name:        name,
		Description: description,
	})
	if err != nil {
		panic(err)
	}
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Admin.Categories())
}
