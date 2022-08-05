
// rpc CreateContact(*models.Contact)(*models.Contact,error)

// can make RPC call client/caller should know the dataformat, method signature

//Protocol Buffers

// .proto


protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/contact.proto

// To update using Bloomrpc

{
  "ID": "6",
  "Data": {
    "fields": {
      "address": {
        "kind": 0,
        "stringValue":"Goa"
      }
    }
  }
}

// For information follow this link https://stackoverflow.com/questions/52966444/is-google-protobuf-struct-proto-the-best-way-to-send-dynamic-json-over-grpc