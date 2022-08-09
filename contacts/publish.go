package main

import (
	"contacts/database"
	"contacts/messaging"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
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
	session, err := database.GetConnection(DSN)
	log.Println("Connecting to the database")
	if err != nil {
		log.Fatalln(err)
	}

	cdb := &database.ContactDB{DB: session}

	for {
		time.Sleep(time.Second * 5)
		conatcts, err := cdb.GetAllByStatus("created")
		if err != nil {
			log.Println(err)
		}
		fmt.Println(conatcts)

		for _, contact := range conatcts {
			msg := new(messaging.Message)
			msg.URL = MESSAGEURL
			msg.Subject = "Contact-Create"
			data, err := contact.ToByte()
			if err != nil {
				log.Println(err)
			}
			msg.Data = data
			if err = msg.Publish(); err != nil {
				log.Println(err)
			} else {
				id := fmt.Sprint(contact.ID)
				_, err := cdb.UpdateBy(id, map[string]interface{}{"Status": "email-sent"})
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
