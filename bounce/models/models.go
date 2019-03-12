package models

import (
	"fmt"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

type Bounces struct {
	time              int
	account_id        int
	sub_accont_id     int
	to_email          string
	from_email        string
	from_domain       string
	injection_ip      int
	remote_ip         int
	bounce_type       string
	bounce_class_name string
	bounce_class_desc string
	classification_id int
	transmission_id   int
	message_id        int
	bounce_reaosn     string
	status            string
}

//
// func (db *DB) PrepareQuery(table_name, chunck_time_interval) stmt {
//
// 	stmt, err := db.PrepareContext("SELECT create_hypertable(table_name=$1, 'time', chunk_time_interval => chunk_time_interval=$2, if_not_exists => TRUE);", table_name, chunk_time_interval)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return stmt
// }
//
// func (db *DB) CreateHypertable(stmt) {
// 	if _, err = db.Exec(stmt); err != nil {
// 		fmt.Println("unable to create schema : ", err)
// 		return
// 	}
// }
//
// func (db *DB) PrepareDataQuery(account_id, table_name) stmt {
// 	stmt, err := db.PrepareContext("SELECT * from table=$1 where $account_id=$2", table_name, account_id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return stmt
// }

func (db *DB) BounceGetData() ([]*Bounces, error) {
	rows, err := db.Query("SELECT * FROM bounces")

	if err != nil {
		fmt.Println("not able to get data", err)
	}
	defer rows.Close()
	bounces := make([]*Bounces, 0)
	for rows.Next() {
		bounce := new(Bounces)
		bounces = append(bounces, bounce)
	}
	return bounces, nil
}
