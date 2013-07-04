package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type Blog struct {
	*revel.Controller
}

func (c Blog) Show(id int, slug string) revel.Result {
	// mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/ironzebra")
	fmt.Println(session)
	fmt.Println(err)
	return c.Render()
}
