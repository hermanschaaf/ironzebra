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
	"bufio"
	"flag"
	"fmt"
	"github.com/hermanschaaf/ironzebra/app/management"
	"github.com/hermanschaaf/revmgo"
	"github.com/robfig/revel"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("\nzebra-admin [OPTIONS] your/app/path [command]\n")
	fmt.Println("List of possible commands:")
	fmt.Println("--------------------------")
	fmt.Println("adduser   [Create a new user]")
	fmt.Println("")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("\n")
}

func loadApp(args []string, mode string) {

	// Find and parse app.conf
	revel.Init(mode, args[0], "")
	revel.LoadMimeConfig()

	revmgo.AppInit()
}

func main() {

	var ip = flag.String("mode", "dev", "Specify which mode to use in app.conf")
	flag.Parse()
	var args = flag.Args()

	if len(args) >= 2 {
		switch args[1] {
		case "adduser":

			fmt.Println("Loading configuration...")
			loadApp(args, *ip)

			var name, username, password string
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Adding a new user")
			fmt.Println("-----------------")

			fmt.Println("What is the user's real name? ")
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("What should the username be? ")
			username, _ = reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Println("Please enter a strong password: ")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)

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
