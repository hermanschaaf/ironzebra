package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Post struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	ShortID   int
	Title     string
	Subtitle  string
	Image     string
	Slug      string
	Body      string
	Category  string
	Tags      []string
	Published bool
	Timestamp time.Time
}

// sequential Id counter for posts
type Counter struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Seq  int
}

func assertSequenceExists(collection *mgo.Collection) {
	c, _ := collection.Find(nil).Count()
	if c == 0 {
		collection.Insert(&Counter{
			Seq: 0})
	}
}

func GetNextSequence(database *mgo.Database) (count int) {
	collection := database.C("postCounter")
	assertSequenceExists(collection)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	result := Counter{}
	collection.Find(nil).Apply(change, &result)
	return result.Seq
}
