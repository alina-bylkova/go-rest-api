## General instructions to create your first program in Go

### Choose a module path and create a go.mod file

go.mod is like package.json in node.js (it contains an inventory of all your dependencies)

go.sum is like package-lock.json in node.js

Run the following commands in the terminal:

```console
$ mkdir <directory-path>

$ cd <directory-path>

$ go mod init <directory-path>

$ touch main.go
```

### Create makefile

makefile allows us to create scripts for our app

```console
$ touch makefile 
```

Inside the file type: `dev: go run main.go`

To run this script in the terminal, type in: 

```console
$ make dev
```

### Install dependencies

To download one specific package, run in the terminal:

```console
$ go get <module-name>
```

To import the package to the file, type in the following to your project file:

```console
import <module-name>
```

To download all required module's dependencies, run in the terminal:

```console
$ go build || go test
```

To list all installed packages, run:

```console
$ go list -m all
```

### Testing

To test specific file, run:

```console
$ go test <test-file-path>
```

**Example:** `go test ./db/slice_db`

To test everything, run:

```console
$ go test ./...
```

To test everything with coverage info, run:

```console
$ go test -cover ./...
```

It will provide coverage for the testing (e.g. the output `coverage: 100%` means that we are covering 100% statements without testing)