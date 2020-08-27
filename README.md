# REST API with GO/Gin

This repo is a simple CRUD API for learning purpose

In this API I have learned features and patterns in Go:

- Using Gin to create a simple API
- Testing HTTP handlers
- Using Docker to containerize the app
- Using Redis for key-value storage

## How can I use it?

**Install**

```sh
$ go get
```

**Usage**

Launch the server

```sh
$ make dev || go run main.go
```

The app will be running at http://localhost:8080/

**Testing**

```sh
$ go test ./... || go test ./db/slice_db
```

## Endpoints

Fetch all numbers

```
GET /numbers
```

Fetch a specific number

```
GET /numbers/{int}
```

Add a number

```
POST /numbers
```

Remove a number

```
DELETE /numbers/{int}
```
