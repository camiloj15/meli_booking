package main

import (
	"log"

	service "profiling/pkg/service"
)

func main() {
	service := service.NewBookService()
	log.Println("Iniciando aplicación...")
	err := service.CreateBook("El señor de los anillos", "J.R.R. Tolkien", 1954, 4.5)
	if err != nil {
		log.Fatal(err)
	}
}
