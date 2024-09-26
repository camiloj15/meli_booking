package repository

import (
	"errors"
	"time"

	"profiling/pkg/model"
)

type BookRepository struct {
	Books []model.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		Books: []model.Book{},
	}
}

func (r *BookRepository) Create(title, author string, year int64, rank float64) error {
	book := &model.Book{
		Title:     title,
		Author:    author,
		Year:      year,
		CreatedAt: time.Now(),
		Ranking:   rank,
	}
	r.Books = append(r.Books, *book)
	return nil
}

func (r *BookRepository) GetBy(title string) (book model.Book, err error) {
	for _, book := range r.Books {
		if book.Title == title {
			return book, nil
		}
	}
	return model.Book{}, errors.New("not Found")
}

func (r *BookRepository) GetTopTenRanked() (books []model.Book, err error) {
	var topTenBooks []model.Book
	copyRepository := r.Books

	numberOfBooks := len(r.Books)
	if numberOfBooks > 0 {
		for i := 0; i < 10; i++ {
			book := copyRepository[i]
			topTenBooks = append(topTenBooks, book)
		}
		for i := 10; i < numberOfBooks; i++ {
			for j := 0; j < 10; j++ {
				if copyRepository[i].Ranking > topTenBooks[j].Ranking {
					oldBook := topTenBooks[j]
					topTenBooks[j] = copyRepository[i]
					copyRepository[i] = oldBook
				}
			}
		}
		for i := 0; i < 9; i++ {
			if topTenBooks[i].Ranking > topTenBooks[i+1].Ranking {
				oldBook := topTenBooks[i+1]
				topTenBooks[i+1] = topTenBooks[i]
				topTenBooks[i] = oldBook
			}
		}
		return topTenBooks, nil
	}

	return nil, errors.New("there are not books in the repository for ranking")
}
