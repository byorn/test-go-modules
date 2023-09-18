package main

import "fmt"
import "github.com/byorn/test-go-modules/util"

func main() {
	fmt.Println("Hello world" + util.ReadFile())

}

func ThisIsSomeMethodToBeCalledFromAnotherRepository() string {
	return "it works"
}
