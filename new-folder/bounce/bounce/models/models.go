package models

import (
	"database/sql"
	"fmt"
)

//
// type Bounces struct {
// 	time              time.Time
// 	account_id        int
// 	sub_account_id    int
// 	to_email          string
// 	from_email        string
// 	from_domain       string
// 	injection_ip      int
// 	remote_ip         int
// 	bounce_type       string
// 	bounce_class_name string
// 	bounce_class_desc string
// 	classification_id int
// 	transmission_id   int
// 	message_id        int
// 	bounce_reaosn     string
// 	status            string
// }

type BouncesReason struct {
	AccountID        int32
	BonuceReason     string
	BounceClassName  string
	ClassificationID int32
	BounceType       string
	BounceClassDesc  string
	Count            int32
}

func BounceGetData(db *sql.DB) ([]*BouncesReason, error) {
	rows, err := db.Query("select bounce_reason,bounce_class_name,classification_id,bounce_type,bounce_class_desc,account_id,count(*) from bounceautodata123 WHERE account_id=$1 GROUP BY bounce_reason,bounce_class_name,classification_id,bounce_type,bounce_class_desc,account_id;")
	if err != nil {
		fmt.Println("not able to get data", err)
	}
	defer rows.Close()
	bounces := make([]*BouncesReason, 0)
	for rows.Next() {
		bounce := new(BouncesReason)
		err := rows.Scan(&bounce.BonuceReason, &bounce.BounceClassName, &bounce.ClassificationID, &bounce.BounceType,
			&bounce.BounceClassDesc, &bounce.AccountID, &bounce.Count)
		if err != nil {
			fmt.Println("Error while retrivng data", err)
		}
		bounces = append(bounces, bounce)
	}
	return bounces, nil
}
