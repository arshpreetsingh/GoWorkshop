package models

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Data struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	value string
}

type UserRepository struct {
	C *mgo.Collection
}

var (
	mongoSession *mgo.Session
)

func AddUser(value string) (string, error) {

	c := getSession().Copy().DB("todo").C("addusers")

	repo := &UserRepository{c}
	return repo.Add(&Data{
		value: value,
	})
}

// create user
func (r *UserRepository) Add(data *Data) (string, error) {
	objID := bson.NewObjectId()
	data.Id = objID
	err := r.C.Insert(&data)
	return data.Id.Hex(), err
}

/*
// Create user
func (r *UserRepository) Create(user *User) (string, error) {
	objID := bson.NewObjectId()
	user.Id = objID
	err := r.C.Insert(&user)
	return user.Id.Hex(), err
}
*/

func getSession() *mgo.Session {
	if mongoSession == nil {
		var err error
		mongoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{os.Getenv("MONGO_HOST")},
			Username: "",
			Password: "",
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.WithError(err).Fatal("could not connecet to Mongo for now")
		}
	}
	return mongoSession

}

/*
func getSession() *mgo.Session {
	if mongoSession == nil {
		var err error
		mongoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{os.Getenv("MONGO_HOST")},
			Username: "",
			Password: "",
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.WithError(err).Fatal("could not connect to mongo")
		}
	}
	return mongoSession
*/
