package main

import (
	"contacts/database"
	grpchandler "contacts/handlers/grpc"
	pb "contacts/protos"
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
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
		flag.StringVar(&PORT, "port", "50081", "--port=50081")
		flag.StringVar(&DSN, "database", `host=localhost user=contactsdb_user password=contactsdb_admin dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--database=valid database connection string")
		flag.Parse()
	}
}

// type server struct {
// 	pb.UnimplementedGreaterServer
// }

// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
// }

func main() {

	log.Println("Contact GRPC Microservice started! about run on PORT", PORT)

	session, err := database.GetConnection(DSN)
	log.Println("Connecting to the database")
	if err != nil {
		log.Fatalln(err)
	}

	cdb := &database.ContactDB{DB: session}
	server := &grpchandler.ContactServer{IContact: cdb}

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	pb.RegisterContactServer(srv, server)

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}
