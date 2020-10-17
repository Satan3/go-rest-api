package main

import(
	"log"
	"github.com/Satan3/go-rest-api/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start()
}
