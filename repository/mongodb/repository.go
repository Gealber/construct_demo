package mongodb

import (
	"github.com/Gealber/construct_demo/accounts"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

const (
	DB_NAME = "ConstructDemo"
)

//MongoRepository structure that represents a MongoDB
//Repository.
type MongoRepository struct {
	URL     string
	Logger  *log.Logger
	Session *mgo.Session
}

//New creates a new MongoRepository and try to stablish the connection
func NewMongoRepository() *MongoRepository {
	url := os.Getenv("MONGO_URL")
	if url == "" {
		url = "localhost"
	}
	logger := log.New(os.Stdout, "[MONGODB]", 0)
	session, err := mgo.Dial(url)
	if err != nil {
		logger.Printf("Unable to comunicate %v\n", err)
		return nil
	}
	return &MongoRepository{
		URL:     url,
		Logger:  logger,
		Session: session,
	}
}

//Find find the user by the id
func (r *MongoRepository) Find(id string) (*accounts.User, error) {
	session := r.Session.Copy()
	defer session.Close()
	coll := session.DB(DB_NAME).C("Users")

	user := &accounts.User{}
	err := coll.FindId(id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//Store will insert a new User in the DB
func (r *MongoRepository) Store(user *accounts.User) error {
	session := r.Session.Copy()
	defer session.Close()
	coll := session.DB(DB_NAME).C("Users")

	err := coll.Insert(user)
	return err
}
