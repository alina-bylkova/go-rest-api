# REST API with GO/Gin

## General instructions to create program in go

### Choose a module path and create a go.mod file

mkdir `<directory-path>`

cd `<directory-path>`

go mod init `<directory-path>`

- go.mod is like package.json in node.js (it contains an inventory of all your dependencies)
- go.sum is like package-lock.json in node.js

touch main.go

touch makefile (it allows us to create scripts for our app)

- inside the file type: `dev: go run main.go`
- to run this command in the terminal, type in `make dev`

### Install dependencies

go get `<module-name>` - to download one specific package

import `<module-name>` to the file - to import package to the file

go build || go test - to download all required module's dependencies

go list -m all - to list all installed packages
