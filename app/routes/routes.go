// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (_ tAdmin) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Logout", args).Url
}

func (_ tAdmin) Edit(
		id int,
		slug string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "slug", slug)
	return revel.MainRouter.Reverse("Admin.Edit", args).Url
}

func (_ tAdmin) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.New", args).Url
}

func (_ tAdmin) SaveNew(
		title string,
		subtitle string,
		category string,
		body string,
		image string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "subtitle", subtitle)
	revel.Unbind(args, "category", category)
	revel.Unbind(args, "body", body)
	revel.Unbind(args, "image", image)
	return revel.MainRouter.Reverse("Admin.SaveNew", args).Url
}

func (_ tAdmin) Save(
		id int,
		title string,
		subtitle string,
		slugString string,
		category string,
		body string,
		publish string,
		image string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "subtitle", subtitle)
	revel.Unbind(args, "slugString", slugString)
	revel.Unbind(args, "category", category)
	revel.Unbind(args, "body", body)
	revel.Unbind(args, "publish", publish)
	revel.Unbind(args, "image", image)
	return revel.MainRouter.Reverse("Admin.Save", args).Url
}

func (_ tAdmin) SaveTags(
		shortid int,
		tags string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "shortid", shortid)
	revel.Unbind(args, "tags", tags)
	return revel.MainRouter.Reverse("Admin.SaveTags", args).Url
}

func (_ tAdmin) RedirectToSlug(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.RedirectToSlug", args).Url
}

func (_ tAdmin) Publish(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.Publish", args).Url
}

func (_ tAdmin) Unpublish(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.Unpublish", args).Url
}

func (_ tAdmin) Categories(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Categories", args).Url
}

func (_ tAdmin) NewCategory(
		name string,
		description string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "description", description)
	return revel.MainRouter.Reverse("Admin.NewCategory", args).Url
}

func (_ tAdmin) DeleteCategory(
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Admin.DeleteCategory", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) LoginPost(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.LoginPost", args).Url
}


type tBlog struct {}
var Blog tBlog


func (_ tBlog) RSS(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.RSS", args).Url
}

func (_ tBlog) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.List", args).Url
}

func (_ tBlog) ListAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Blog.ListAll", args).Url
}

func (_ tBlog) ListCategory(
		category string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Blog.ListCategory", args).Url
}

func (_ tBlog) ListTag(
		category string,
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("Blog.ListTag", args).Url
}

func (_ tBlog) Show(
		category string,
		id int,
		slugString string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "slugString", slugString)
	return revel.MainRouter.Reverse("Blog.Show", args).Url
}

func (_ tBlog) RedirectToSlug(
		category string,
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Blog.RedirectToSlug", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


