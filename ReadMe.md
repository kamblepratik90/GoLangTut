`go mod init hello`

`go run main.go`

// to create executables for diff OS

`GOOS="linux" go build`

`GOOS="windows" go build`

`go build`

--------------------------------------------

23 - tut

1. Create folder
2. Create main.go file
3. run - `go mod init github.com/kamblepratik90/mygolangmodules`
4. run - `go get -u github.com/gorilla/mux`
5. after line 4, the respective module will be downloaded and cached in go dir on machine. to verify this we can refer `go env` and navigate to 
GOPATH='/Users/pratikkamble/go'
and then /Users/pratikkamble/go/pkg/mod/cache/download/github.com/gorilla/mux
5. Write the code for routes and server. mention port number.
6. Visit -> localhost:3000
7. run - `go mod tidy` - to make direct dependency for imported modules. And remove unused dependency (expensive operation)
8. run `go mod verify` - to verify all the modules available inside go.sum file
9. run `go list -m all` - to get list of all dependent packages
10. run `go list -m -versions github.com/gorilla/mux` - to get all available versions for this module
11. run `go mod why github.com/gorilla/mux` - to get what dependent on what
12. run `go mod graph` - to check dependency flow

IMP ->

13. run `go mod vendor` - it create a folder name 'vendor' it is similar to 'node_modules' folder, it carries the all the dependencies instead keeping it in system cache.

14. run `go run -mod=vendor main.go` - to user vendor mod dependency.

