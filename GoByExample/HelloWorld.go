// Our first program will print the classic "hello world" message. Here`s the full source code.
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

// To run the program, put the code in hello-world.go and use go run.
// Sometimes we`ll want to build our programs into binaries.
// we can do no this using go build.

// We can then execute the built binary directly.

// Now that we can run and build basic Go programs, let`s
// learn more about the language.

// >> go run HelloWorld.go
// >> hello world
// >> go build HelloWorld.go
// >> ls
// >> HelloWorld HelloWorld.go
// >> ./HelloWorld
// >> hello world