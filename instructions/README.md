## General instructions to create your first program in Go

### Choose a module path and create a go.mod file

Run the following commands in the terminal:

```console
$ mkdir `<directory-path>`

$ cd `<directory-path>`

$ go mod init `<directory-path>`

- go.mod is like package.json in node.js (it contains an inventory of all your dependencies)
- go.sum is like package-lock.json in node.js

$ touch main.go

$ touch makefile 

- it allows us to create scripts for our app
- inside the file type: `dev: go run main.go`
- to run this command in the terminal, type in `make dev`
```

### Install dependencies

Run the following commands in the terminal:

go get `<module-name>` - to download one specific package

- import `<module-name>` to the file - to import package to the file

go build || go test - to download all required module's dependencies

go list -m all - to list all installed packages

### Testing

Run the following commands in the terminal:

go test `<test-file-path`

example: go test ./db/slice_db

go test ./... - to test everything

go test -cover ./... - it will provide coverage for the testing (e.g. the output coverage: 100% means that we are covering 100% statements without testing)