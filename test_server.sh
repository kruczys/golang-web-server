#! /bin/bash

echo "---------------Testing GET /posts---------------"
curl -s http://localhost:8080/posts | jq .
echo ""

echo "---------------Testing POST /posts---------------"
curl -s -X POST -H "Content-Type: application/json" -d '{
    "date": "2024-05-04",
    "country": "PL",
    "name": "Jan Kowalski",
    "activity": "Fishing",
    "age": "47",
    "injury": "Fatal"
}' http://localhost:8080/posts | jq .
echo ""

echo "---------------Testing GET /posts/0---------------"
curl -s http://localhost:8080/posts/0 | jq .
echo ""

echo "---------------Testing DELETE /posts/0---------------"
curl -s -X DELETE http://localhost:8080/posts/0
echo ""

echo "---------------Testing GET /posts after DELETE---------------"
curl -s http://localhost:8080/posts | jq .
echo ""

