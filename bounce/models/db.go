package models

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Datastore interface {
	BounceGetData() ([]*Bounces, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.WithError(err).Fatal("could not connect to Database")
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.WithError(err).Fatal("could not Ping to Database")
		return nil, err
	}
	return &DB{db}, nil
}
