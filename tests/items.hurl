# List items - should have at least the 2 seeded items
GET http://localhost:8080/items
HTTP 200
[Asserts]
jsonpath "$" count >= 2

# Create a new item
POST http://localhost:8080/items
Content-Type: application/json
{"name": "cherry"}
HTTP 201
[Captures]
new_id: jsonpath "$.id"
[Asserts]
jsonpath "$.name" == "cherry"

# Get the newly created item
GET http://localhost:8080/items/{{new_id}}
HTTP 200
[Asserts]
jsonpath "$.name" == "cherry"

# Delete it
DELETE http://localhost:8080/items/{{new_id}}
HTTP 204

# Confirm it's gone
GET http://localhost:8080/items/{{new_id}}
HTTP 404
