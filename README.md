# Algorithms

This is a bucket to practice algorithms using golang

for creating a go module run `go mod init ${module_name}`

for creating a test add a file named __\\_test__ after the file you want to test

i.e hello.go -> hello_test.go

for running a test inside a module run `go test`

__This is the way of renaming default imports__
import quoteV3 "rsc.io/quote/v3"

for cleaning dependencies `go mod tidy`
for adding the module to the work directory (inside the specific module folder) `go work use .`

__run__

if one wants to execute a program, it need to change the package to main
build the package typing `go build *.go` and then execute the
generated binary passing the arguments using  `go run *.exe`