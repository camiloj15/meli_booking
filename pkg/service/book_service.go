package service

import (
	"profiling/pkg/repository"
	"profiling/pkg/model"
)

type BookService struct {
	repositorio *repository.BookRepository
}

func NewBookService() *BookService {
	return &BookService{
		repositorio: repository.NewBookRepository(),
	}
}

func (s *BookService) CrearLibro(titulo, autor string, año int64) error {
	err := s.repositorio.CrearLibro(titulo, autor, año)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetLibros() []model.Book {
	return s.repositorio.Libros
}

func (s *BookService) GetBy(titulo string) (model.Book, error) {
	return s.repositorio.GetBy(titulo)
}