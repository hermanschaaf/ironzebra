package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Post struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	ShortID   int
	Title     string
	Subtitle  string
	Slug      string
	Body      string
	Timestamp time.Time
}
