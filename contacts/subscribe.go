package main

import (
	"contacts/messaging"
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	PORT       string
	DSN        string
	MESSAGEURL string
)

func init() {
	PORT = os.Getenv("PORT")
	DSN = "host=localhost user=contactsdb_user password=contactsdb_admin dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	MESSAGEURL = "nats://127.0.0.1:4222"
}

func init() {
	if PORT == "" {
		flag.StringVar(&PORT, "port", "50080", "--port=50080")
		flag.StringVar(&DSN, "database", `host=localhost user=contactsdb_user password=contactsdb_admin dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--database=valid database connection string")
		flag.StringVar(&MESSAGEURL, "msgbroker", `nats://127.0.0.1:4222`, "--msgbroker=valid message broker connection string")

		flag.Parse()
	}
}
func main() {
	msg := new(messaging.Message)
	msg.URL = "nats://localhost:4222"
	msg.Subject = "contact.create"
	//for {
	//time.Sleep(time.Second * 1)
	fmt.Println("Subscribing---")
	err := msg.Subscribe()
	if err != nil {
		fmt.Println(err)
	}
	//}
	runtime.Goexit()
}

/*Publish --> Topic/Subject
Process --> Subscribe a topic, process it and publish to another topic
Subscribe --> Reads from the Topic do some stuff*/
