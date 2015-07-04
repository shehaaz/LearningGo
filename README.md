# LearningGo
Examples using GoLang

Installing on MAC

`//If you already have an old GO`

```sh 
$ brew uninstall go hg
```

`$ brew install go hg`

`//create a local GO workspace`
`$ mkdir $HOME/go`

`$ export GOPATH=$HOME/go`

`$ export PATH=GOPATH/bin:$PATH`

`//This is a convention in GO`
`$ mkdir -p $GOPATH/src/github.com/[GITHUB_USERNAME]/[PROJECT_NAME]`


`$ cd $GOPATH/src/github.com/[GITHUB_USERNAME]/[PROJECT_NAME]`

`$ touch hello.go`

```go
package main
import "fmt"
func main() {
    fmt.Printf("Hello, world.\n")
}```

`//this command will create $GOPATH/bin folder and put in the compiled "hello"`
`$ go install`

`$ hello`
`>Hello, world.`
