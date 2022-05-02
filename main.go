package main

import (
	"github.com/willie-lin/Go_With_Test/dependencyinjection"
	"os"
)

func main() {
	dependencyinjection.Greet(os.Stdout, "Elodie")

}
