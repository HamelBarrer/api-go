package main

import (
	"log"

	"github.com/HamelBarrer/api-go/connections"
	"github.com/HamelBarrer/api-go/handlers"
)

func main() {
	if !connections.ConnectionPing() {
		log.Fatal("sin conexion a mongodb")
		return
	}

	handlers.Handler()
}
