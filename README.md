# REST API with GO/Gin

This repo is a simple CRUD API for learning purpose

In this API I have learned features and patterns in Go:

- Using Gin to create a simple API
- Using in-memory database to store an array of numbers 
- Testing HTTP handlers
- Using Docker to containerize the app
- Implementing Redis for key-value storage

## How can I use it?

### For local development follow these steps:

**Install**

```sh
$ go mod download
```

**Usage**

Launch the server

```sh
$ make dev
```

OR

```sh
$ go run main.go
```

The app will be running at http://localhost:8080/

**Testing**

```sh
$ go test ./...
```

OR

```sh
$ go test ./db/slice_db && go test ./db/map_db
```

## Endpoints

Fetch all numbers

```
GET /numbers
```

Fetch a specific number

```
GET /numbers/:int
```

Add a number

```
POST /numbers
```

Remove a number

```
DELETE /numbers/:int
```


## Build docker container with the app

### In case if you need to build the go app for linux, uncomment first option in Docker file and run this command in the terminal:
`env GOOS=linux GOARCH=amd64 go build`

### Build docker image
`docker build -t rest-api:1.0 .`

where `rest-api` is image name and `1.0` is a tag

### Start containers in detached mode
`docker-compose up -d`

## Helpful commands 

### Run a command inside a container
If you want to connect to redis cli run:

`docker exec -it {container_name} {command}`

To connect to redis-cli:

`docker exec -it rd redis-cli` 
