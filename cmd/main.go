package main

import (
	"fmt"

	"github.com/vishal2911/school/server"
)

func main() {
	server := server.Server{}
	server.NewServer()
	fmt.Println("Pratyush Running Server......")
}