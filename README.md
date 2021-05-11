db-test
-------

This is a very simple test to check what could be the cost of
creating a db connection per request instead of using a connection pool.

### Requirements

* docker-compose
* go lang 1.16

### Run & build

Restore dependencies
```shell
go mod tidy
```

Start db instance
```shell
docker-compose up -d
```

Execute the app
```shell
go run cmd/main.go
```

Fetch using the connection pool
```
GET http://localhost:5000/v1/dummies
```

Fetch creating a connection per request
```
GET http://localhost:5000/v2/dummies
```