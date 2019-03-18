package dbWrite

import (
    "fmt"
    "encoding/json"
    // "string"
    // "os"
    "log"
    "relay/reports/ip_reports/grpc/server/database"
    "github.com/streadway/amqp"
)

// type RabbitMQData struct {
// 	AccountID  string `json:"account_id"`
// 	From       string `json:"from"`
// 	Rcpt       string `json:"rcpt"`
// 	Domain     string `json:"domain"`
// 	Event      string `json:"event"`
// 	BounceType string `json:"bounce_type"`
// 	Diagnostic string `json:"diagnostic"`
// 	Retry      int    `json:"retry"`
// 	IP         string `json:"ip"`
// }

func init() {
    database.Create_connection()
}

func InsertData(a amqp.Delivery){
    var data = map[string]string{} 
    fmt.Println("db connected...", a) 
    json.Unmarshal(a.Body, &data)
	// log.Printf(" [x] %s", a.Body)  

    query := "INSERT INTO reports(recorded_at,account_id,sub_account_id,to_email,from_email,to_domain,from_domain,injection_ip,remote_ip,ip_pool,event) VALUES (now()," + data["account_id"] + ", 1,'" + data["rcpt"] + "','" + data["from"] + "','to_domain','" + data["domain"] + "','" + data["ip"] + "','0.0.0.127','ip_pool','" + data["event"]+ "');"
    fmt.Println(query)

    _, err := database.DB.Query(query)
    if err != nil {
        log.Fatalf("Error ocurred while select query: %v", err)
    }

    fmt.Println("Data Inserted Successfully")
}
