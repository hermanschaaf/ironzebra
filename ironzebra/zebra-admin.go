package main

// This is an admin tool for managing Zebra blogs
// It's a work in progress - there is no proper error-handling,
// so if it just explodes, you'll probably have to trace the code
// to see what's going wrong. 
//
// If I have to guess though, it's probably that some settings
// (like db.name and revmgo.dial) are not set in app.conf, or that
// the database key you are trying to create already exists.

import (
	"fmt"
	"github.com/hermanschaaf/ironzebra/app/management"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"os"
)

func printHelp() {
	fmt.Println("\nzebra-admin your/app/path [command]\n")
	fmt.Println("List of possible commands:")
	fmt.Println("--------------------------")
	fmt.Println("adduser   [Create a new user]")
	fmt.Println("\n")
}

func loadApp(args []string) {
	mode := "dev"

	// Find and parse app.conf
	revel.Init(mode, args[0], "")
	revel.LoadMimeConfig()

	revmgo.AppInit()
}

func main() {
	args := os.Args[1:]

	if len(args) >= 2 {
		switch args[1] {
		case "adduser":

			fmt.Println("Loading configuration...")
			loadApp(args)

			var name, username, password string

			fmt.Println("Adding a new user")
			fmt.Println("-----------------")

			// fmt.Println("What is the user's real name? ")
			// fmt.Scanf("%s", &name)

			// fmt.Println("What should the username be? ")
			// fmt.Scanf("%s", &username)

			// fmt.Println("Please enter a strong password: ")
			// fmt.Scanf("%s", &password)
			name = "Herman Schaaf"
			username = "hermanschaaf"
			password = "bakabakabaka"

			session := revmgo.Session.New()
			users := session.DB(revmgo.Database).C("users")
			fmt.Println(name, username, password)

			management.AddUser(users, name, username, password)
		default:
			printHelp()
		}
	} else {
		printHelp()
	}
}
