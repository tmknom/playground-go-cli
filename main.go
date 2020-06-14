package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	envPort, _ := strconv.Atoi(os.Getenv("EXAMPLE_PORT"))
	var port int

	flag.IntVar(&port, "p", envPort, "port to use (short)")
	flag.IntVar(&port, "port", envPort, "port to use")
	flag.Parse()

	fmt.Printf("port: %d\n", port)
}
