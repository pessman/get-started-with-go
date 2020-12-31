# Go Hello World

## Run code

```bash
go run hello.go
```

## Using external module

[Link](https://golang.org/doc/tutorial/call-module-code)

```bash
go build
./hello
```

## Compile and install the application

[Link](https://golang.org/doc/tutorial/compile-install)

```bash
go list -f '{{.Target}}'
export PATH=$PATH:/path/to/your/install/directory
go install
hello
```