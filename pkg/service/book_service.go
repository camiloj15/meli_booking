package service

import (
	"profiling/pkg/model"
	"profiling/pkg/repository"
)

type BookService struct {
	bookRepository *repository.BookRepository
}

func NewBookService() *BookService {
	return &BookService{
		bookRepository: repository.NewBookRepository(),
	}
}

func (s *BookService) CreateBook(title, author string, year int64, rank float64) error {
	err := s.bookRepository.Create(title, author, year, rank)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetBooks() []model.Book {
	return s.bookRepository.Books
}

func (s *BookService) GetBy(title string) (model.Book, error) {
	return s.bookRepository.GetBy(title)
}

func (s *BookService) GetTopTenRanked() ([]model.Book, error) {
	return s.bookRepository.GetTopTenRanked()
}
