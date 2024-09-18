package main_test

import (
	"testing"
	"fmt"

	"profiling/pkg/service"
)

func TestCrearLibros(t *testing.T) {
	service := service.NewBookService()
	length := 100000
	for i := 0; i < length; i++ { // Crea 10,000 libros
		err := service.CrearLibro(fmt.Sprintf("Título %d", i), fmt.Sprintf("Autor %d", i), int64(i))
		if err != nil {
			t.Errorf("Error creando libro: %v", err)
			return
		}
	}
	// Verificar que se hayan creado correctamente
	libros := service.GetLibros()
	if len(libros) != length {
		t.Errorf("Se esperaban 10,000 libros, pero se encontraron %d", len(libros))
	}

	for i := 0; i < length; i++ {
		_, err := service.GetBy(fmt.Sprintf("Título %d", i))
		if err != nil {
			t.Errorf("Error obteniendo libro: %v", err)
			return
		}
	}
}
