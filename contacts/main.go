package main1

import (
	"contacts/database"
	"contacts/handlers"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

func init() {
	PORT = os.Getenv("PORT")
	DSN = "host=localhost user=contactsdb_user password=contactsdb_admin dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
}

func init() {
	if PORT == "" {
		flag.StringVar(&PORT, "port", "50080", "--port=50080")
		flag.StringVar(&DSN, "database", `host=localhost user=contactsdb_user password=contactsdb_admin dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--database=valid database connection string")
		flag.Parse()
	}
}

func main() {

	log.Println("Contact Microservice started! about run on PORT", PORT)
	session, err := database.GetConnection(DSN)
	log.Println("Connecting to the database")
	if err != nil {
		log.Fatalln(err)
	}

	cdb := &database.ContactDB{DB: session}

	//cdb := &filedb.ContactFDB{DB: "contacts.fdb"}

	// contact := &models.Contact{Name: "Jiten", Email: "Jitenp@outlook.com", Address: "Rajajinagar, Bangalore", MoreInfo: "Testing contact", ContactNo: "9618558500", Status: "Created", LastModified: fmt.Sprint(time.Now().Unix())}

	// c, err := cdb.Create(contact)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(c)
	// }

	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	r.GET("/health", handlers.Health())

	r.GET("/greet", handlers.Greet("Hello World!"))

	// contacts api

	chandler := handlers.Contact{IContact: cdb}

	v1contactGroup := r.Group("/v1/contact")
	v1contactGroup.POST("/create", chandler.Create())
	v1contactGroup.GET("/get/:id", chandler.GetBy())
	v1contactGroup.PUT("/update/:id", chandler.UpdateBy())
	v1contactGroup.DELETE("/delete/:id", chandler.DeleteBy())

	log.Fatal(r.Run(":" + PORT)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// :50050 what is the ip address?
// 0.0.0.0 != 127.0.0.1
// 127.0.0.1
// 192.168.0.4
// 10.45.0.4
// 192.168.0.1/24 -- 192.168.0.1 --- 192.168.0.255
// netwrok interfaces
// LO,ETH0,WSLP,
// If your computer is connectec through a model to a wireless router --> There must be an interface
// loopback--> connecting to your own system.
// eithernet interface

// A simple CRUD operations on Contacts using Postgres
// Have Mongodb as another database and use interfaces concept and use kind of dependency injection
// Implement context (We will discuss more about context)
// gomock / How to write tests writing few tests(2-3 tests)
// implement NATS as a message broker.
// GRPC based application similar application.
// go escape analysis and profiling to check how stuff works
// check and write ginko tests (introduction and one or more tests)
// Dockerizing the whole application. Use Docker-Compose time permits we use Kubernetes.
