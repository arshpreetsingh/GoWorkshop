package models

import (
"os"
"time"
mgo "gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type Data struct{
 value string
}

type UserRepository struct {
  C *mgo.Collection
}

type

var mongosession *mgo.Session

func AddUser(value string) string ,error{

  c := getSession().Copy().DB("todo").C("addusers")

  repo := &UserRepository{c}
  return repo.Add(&Data{
    value: value,
  })
}

// create user
func (r *UserRepository) Add(data *Data) (string, error){
  err := r.C.Insert(&data)
	return err
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

func getSession() *mgo.Session{
  if mongoSession==nil {
    var err error
    mongoSession,err=mgo.DialWithInfo($mgo.DialInfo{
      Addrs: []string{os.Getenv("MONGO_HOST")},
      Username: "",
      Password: "",
      Timeout: 60*time.Second,

    })
    if err!=nil{
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

func (r *UserRepository)Add(value *Data) string,error{


}
