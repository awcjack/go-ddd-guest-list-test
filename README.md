# Guest List system

A guest list system that implemented following DDD philosophy using MySQL database as persistent storage. This system can be accessed via RESTful API at http://localhost:3000 after start up.

## Pre-requisite
- Docker (Or docker alternative with docker-compose)
- cmake
- amd64 based architecture instance (For MySQL 5.7, Possible to use v8/arm64 instance using qemu)

## How to start
`make docker-up`  
After running this command, a MySQL database is started at port 3306 and the go application is started at port 3000.

## How to test
unit test: `make unit-test`  
integration test: `make integration` (a test MySQL database will be started)  
end-to-end test: `make docker-up` and then test with your favourite HTTP client (e.g. common GUI http client Postman or common DOS http client curl/wget)  

## Why
DDD: Provide flexibility to switch to different storage layer (infrastructure layer that can provide memory, MySQL, PostgreSQL, MongDB, Redis, etc. database), different server layer (interface layer that can provide HTTP, gRPC, etc server implementation)  
Environment Config: reading config from environment is easiest way for deploying application to Docker/Kubernetes  
OpenAPI: openapi provide a easy way to building server and client based on the same specification (like using the same proto in gRPC)

## Improvment  
Using UUID/GUID as unique identifier in Table  
Using other unique identifier for guest as index due to  different people may have same name and also using string as index is a bit bad. The guest is expected to have other info like user (if public event) or employee (if the event is hosted by company) so using userId or employeeId as primary key is preferred  
Better business and coding logic separation  
More log provided in different log level  
Higher test coverage  
CI/CD with security scanning (static scan, gitleak, etc.) test stage, deployment in production environment  
Provide different set of sample data for integration test

## API provided in this application
**Open API specification is provided in docs/openapi.yaml**

API description below is copied from the README provided by the assignment
### Add table

```
POST /tables
body: 
{
    "capacity": 10
}
response: 
{
    "id": 2,
    "capacity": 10
}
```

### Add a guest to the guestlist

If there is insufficient space at the specified table, then an error should be thrown.

```
POST /guest_list/name
body: 
{
    "table": int,
    "accompanying_guests": int
}
response: 
{
    "name": "string"
}
```

### Get the guest list

```
GET /guest_list
response: 
{
    "guests": [
        {
            "name": "string",
            "table": int,
            "accompanying_guests": int
        }, ...
    ]
}
```

### Guest Arrives

A guest may arrive with an entourage that is not the size indicated at the guest list.
If the table is expected to have space for the extras, allow them to come. Otherwise, this method should throw an error.

```
PUT /guests/name
body:
{
    "accompanying_guests": int
}
response:
{
    "name": "string"
}
```

### Guest Leaves

When a guest leaves, all their accompanying guests leave as well.

```
DELETE /guests/name
response code: 204
```

### Get arrived guests

```
GET /guests
response: 
{
    "guests": [
        {
            "name": "string",
            "accompanying_guests": int,
            "time_arrived": "string"
        }
    ]
}
```

### Count number of empty seats

```
GET /seats_empty
response:
{
    "seats_empty": int
}
```