package management

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/hermanschaaf/ironzebra/app/models"
	"labix.org/v2/mgo"
)

func AddUser(collection *mgo.Collection, name, username, password string) {
	// Index
	index := mgo.Index{
		Key:        []string{"username", "email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)

	// Insert Dataz
	err = collection.Insert(&models.User{Name: name, Username: username, Password: bcryptPassword})

	if err != nil {
		panic(err)
	}
}
