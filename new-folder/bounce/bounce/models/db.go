package models

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// var schema = `
// CREATE TABLE IF NOT EXISTS bouncetest (
// 	time              INT					NOT NULL,
// 	account_id        INT					NOT NULL,
// 	sub_account_id     INT					NOT NULL,
// 	to_email          TEXT			 NULL,
// 	from_email        TEXT			 NULL,
// 	from_domain       TEXT 			 NULL,
// 	injection_ip      DOUBLE PRECISION					 NULL,
// 	remote_ip         DOUBLE PRECISION					 NULL,
// 	bounce_type       TEXT			 NULL,
// 	bounce_class_name TEXT			 NULL,
// 	bounce_class_desc TEXT			 NULL,
// 	classification_id INT 				 NULL,
// 	transmission_id   INT					 NULL,
// 	message_id        INT					 NULL,
// 	bounce_reaosn     TEXT			 NULL,
// 	status            TEXT			 NULL
// 	);`

// create init function here to make entries into DB as well
// var schema2 = `INSERT INTO bounces(time,account_id,sub_account_id,to_email,from_email,from_domain,
// 	injection_ip,remote_ip,bounce_type,bounce_class_name,bounce_class_desc,classification_id,
// 	transmission_id,message_id,bounce_reaosn,status)
// VALUES (NOW(), 4325234,435643,"arshpreetjfg",
// "maropost.com",1242134215,2352345,"hard","ust new bounce",
// "this is desc of bunce class as well",234234532,2353425334,345345,"bounce reson invaid","status bounce");`
// var schema2 = `INSERT INTO bouncetest(time,acount_id,sub_account_id)VALUES(12231,2134,213423)
// `

// func init() {
// 	var err error
// 	// test without port thing
// 	connStr := "user=test dbname=test password=test host=0.0.0.0 sslmode=disable port=5000"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.WithError(err).Error("could not connecet to Database")
// 	}
// 	if pingerr := db.Ping(); pingerr != nil {
// 		fmt.Println("unable to ping DB  : " + pingerr.Error())
// 		return
// 	}
// 	log.Info("Successfully connected! to Timecale DB")
//
// 	return
// }

func NewDB() (*sql.DB, error) {
	//db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", "test", "password", "dev", os.Getenv("TIMESCALE_HOST"), "5432"))
	connStr := "user=test dbname=test password=test host=0.0.0.0 sslmode=disable port=5000"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.WithError(err).Fatal("could not connect to Database")
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.WithError(err).Fatal("could not Ping to Database")
		return nil, err
	}
	return db, nil
}
