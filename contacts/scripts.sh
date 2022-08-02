#! /bin/bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"xyz","email":"xyz@gmail.com","address":"bangalore,rajajinagar","contactNo":"9618558500","moreInfo":"Test through curl script"}' \
  http://localhost:50080/v1/contact/create