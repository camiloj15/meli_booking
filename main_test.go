package main_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"profiling/pkg/service"
)

func TestCrearLibros(t *testing.T) {
	bookService := service.NewBookService()
	length := 10000
	for i := 0; i < length; i++ {
		err := bookService.CreateBook(fmt.Sprintf("Título %d", i), fmt.Sprintf("Author %d", i), int64(i), randFloats(1, 5))
		if err != nil {
			t.Errorf("Error creando libro: %v", err)
			return
		}
	}

	for i := 0; i < length; i++ {
		_, err := bookService.GetBy(fmt.Sprintf("Título %d", i))
		if err != nil {
			t.Errorf("Error obteniendo libro: %v", err)
			return
		}
	}

	topTenBooks, errTopTen := bookService.GetTopTenRanked()
	if errTopTen == nil {
		t.Log("This is the top ten")
		for _, book := range topTenBooks {
			t.Log(book)
		}
	}
}

func randFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
