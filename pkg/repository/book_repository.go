package repository

import (
	"errors"
	"sync"
	"time"

	"profiling/pkg/model"
)

type BookRepository struct {
	Libros []model.Book
	mu     sync.RWMutex
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		Libros: []model.Book{},
	}
}

func (r *BookRepository) CrearLibro(titulo, autor string, año int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	libro := &model.Book{
		Titulo:   titulo,
		Autor:    autor,
		Año:      año,
		CreadoEn: time.Now(),
	}
	r.Libros = append(r.Libros, *libro)
	return nil
}

func (r *BookRepository) GetBy(titulo string) (book model.Book, err error) {
	for _, book := range r.Libros {
		if(book.Titulo == titulo){
			return book, nil
		}
	}
	return model.Book{}, errors.New("Not Found")
}
