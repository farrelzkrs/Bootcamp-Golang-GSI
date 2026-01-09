package main

import (
	"flag"
	"fmt"
)

func main() {
	var hostname *string = flag.String("host", "localhost", "")
	var port *int = flag.Int("port", 8080, "")
	var password *string = flag.String("password", "", "")

	flag.Parse()

	fmt.Println("Hostname:", *hostname)
	fmt.Println("Port:", *port)
	fmt.Println("Password:", *password)
}