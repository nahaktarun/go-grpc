package main

import (
	"log"

	proto "github.com/nahaktarun/grpc-module-1/proto"
)

func main() {

	person := proto.Person{Name: "Tarun"}

	log.Println(person.GetName())
}
