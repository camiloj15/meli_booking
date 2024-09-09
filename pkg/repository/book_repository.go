package repository

import (
	"sync"
	"time"

	"profiling/pkg/model"
)

type BookRepository struct {
	Libros map[string]*model.Book
	mu     sync.RWMutex
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		Libros: make(map[string]*model.Book),
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
	r.Libros[titulo] = libro
	return nil
}
