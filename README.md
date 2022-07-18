- To run

```go run main.go```

- To build

```go build main.go```

- To build to a specific name

```go build -o app main.go```

- List of the target GOOS/GOARCH

```go tool dist list```

- To build on other targeted machines

```GOOS=windows GOARCH=amd64 go build -o app.exe main.go```

- To install go binaries in GOBIN path

```go install main.go```
or
```go install github.com/XXXXXXXX```

- To create go module

```go mod init demo```

- The name of the module is the root path for all user defined packages in the project