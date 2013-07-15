package controllers

import (
	"bitbucket.org/gosimple/slug"
	"code.google.com/p/go.crypto/bcrypt"
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

func addPost(database *mgo.Database, collection *mgo.Collection, title, subtitle, slug, body string) (post models.Post) {
	// Index
	index := mgo.Index{
		Key:        []string{"shortid", "timestamp", "title", "tags"},
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

func savePost(collection *mgo.Collection, shortID int, title, subtitle, slugString, body string) (post models.Post) {
	// Update Dataz
	err := collection.Update(bson.M{"shortid": shortID}, bson.M{
		"$set": bson.M{
			"title":    title,
			"subtitle": subtitle,
			"slug":     slugString,
			"body":     body,
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
	username, loggedIn := c.Session["username"]
	name, _ := c.Session["name"]
	if loggedIn == true {
		collection := c.Database.C("posts")
		postList := []models.Post{}
		iter := collection.Find(nil).Iter()
		iter.All(&postList)
		return c.Render(username, name, loggedIn, postList)
	}
	return c.Render(loggedIn)
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
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err == nil {
			c.Session["username"] = username
			c.Session["name"] = user.Name
			c.Session["role"] = "admin"
			c.Flash.Success("Welcome, " + user.Name)
			return c.Redirect(routes.Admin.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Admin.Index())
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
		return c.Redirect(routes.Blog.Show(id, result.Slug))
	}

	// TODO: this is a bit of a hack
	if result.Slug == "" {
		return c.Redirect(routes.Admin.Index())
	}

	newPost := !result.Published

	return c.Render(result, newPost)
}

func (c Admin) New() revel.Result {
	return c.Render()
}

func validatePost(c Admin, title, body, slugString string) {
	c.Validation.Required(title).Message("A title is required")
	c.Validation.Required(body).Message("You probably want some text in your post, no?")
	c.Validation.Required(slugString).Message("You need a slug...")
}

func (c Admin) SaveNew(title, subtitle, body string) revel.Result {
	slugString := slug.Make(title)
	validatePost(c, title, body, slugString)
	collection := c.Database.C("posts")
	result := addPost(c.Database, collection, title, subtitle, slugString, body)
	return c.Redirect(routes.Blog.Show(result.ShortID, result.Slug))
}

func (c Admin) Save(id int, title, subtitle, slugString, body, publish string) revel.Result {
	validatePost(c, title, body, slugString)
	collection := c.Database.C("posts")
	if slugString == "" {
		slugString = slug.Make(title)
	}
	result := savePost(collection, id, title, subtitle, slugString, body)
	return c.Redirect(routes.Blog.Show(result.ShortID, result.Slug))
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
	return c.Redirect(routes.Blog.Show(result.ShortID, result.Slug))
}

func (c Admin) Unpublish(id int) revel.Result {
	/* Unpublish the post */
	collection := c.Database.C("posts")
	collection.Update(bson.M{"shortid": id}, bson.M{"$set": bson.M{"published": false}})

	result := models.Post{}
	collection.Find(bson.M{"shortid": id}).One(&result)
	return c.Redirect(routes.Blog.Show(result.ShortID, result.Slug))
}

func (c Admin) AddImages() revel.Result {
	return c.Render()
}
