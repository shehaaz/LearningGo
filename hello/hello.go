package main

/*
Notice we imported the string file that contained the Reverse method
*/
import (
	"fmt"

	"github.com/shehaaz/LearningGo/string"
)

func main() {
	fmt.Println("Hello, new gopher")
	fmt.Println(string.Reverse("Hello, new gopher"))
}

/*

To run the program, put the code in hello-world.go and use go run.
$ go run hello.go
"Hello, new gopher"
rehpog wen ,olleH

Sometimes we’ll want to build our programs into binaries. We can do this using go build.
$ go build hello.go
$ ls
hello	hello.go
We can then execute the built binary directly.
$ ./hello
"Hello, new gopher"
rehpog wen ,olleH

$ go install
//This command builds the hello.go program into ~/go/bin/
//now you can simply run it using:
//remember to add bin directory to the system PATH
$ export PATH=$HOME/gocode/bin:$PATH
$ hello
"Hello, new gopher"
rehpog wen ,olleH

Now that we can run and build basic Go programs, let’s learn more about the language.


*/
