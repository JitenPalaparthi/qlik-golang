#! /bin/bash

#POST
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"xyz","email":"xyz@gmail.com","address":"bangalore,rajajinagar","contactNo":"9618558500","moreInfo":"Test through curl script"}' \
  http://localhost:50080/v1/contact/create

#GET
curl http://localhost:50080/v1/contact/get/1

#DELETE
 curl -X "DELETE" http://localhost:50080/v1/contact/delete/2

#PUT
curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"address":"bangalore,Mahalakshmi Layout"}' \
  http://localhost:50080/v1/contact/update/1


docker run -d --name pg -p 5432:5432 -e POSTGRES_PASSWORD=contactsdb_admin -e POSTGRES_USER=contactsdb_user -e POSTGRES_DB=contactsdb postgres
docker run -d --name nats1 -p 4222:4222 -p 8222:8222 nats

docker run -d --name c1 -p 50080:50080 -e DSN="host=172.17.0.3 user=contactsdb_user password=contactsdb_admin dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai" -e MESSAGEURL="nats://172.17.0.2:4222" jpalaparthi/contacts

docker build . -t jpalaparthi/contacts
# container image
# database image
# message broker image

# 3-kubernets deployments
# 3 services