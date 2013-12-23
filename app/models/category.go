package models

import (
	"labix.org/v2/mgo/bson"
)

type Category struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string
}
