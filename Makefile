# The .DEFAULT_GOAL defines which target is run when no target is specified. In this case, the default is the build target. So when you run `make` (without specifying a target), you will see output like this:
# $ make
# go clean ./...
# go fmt ./...
# go vet ./...
# go build
.DEFAULT_GOAL := build

# The .PHONY line keeps make from getting confused if a directory or file in your project has the same name as one of the listed targets.
.PHONY: dev clean fmt vet build

dev:
	go run main.go

# Each possible operation is called a target and the following definitions are the target definitions. 
# The word before the colon (:) is the name of the target. 
# Any words after the target (like `vet` in the line `build: vet`) are the other targets that must be run before the specified target runs. 
# The tasks that are performed by the target are on the indented lines after the target.
# Using ./... tells a Go tool to apply the command to all the files in the current directory and all subdirectories.
clean:
	go clean ./...

fmt: clean
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build
