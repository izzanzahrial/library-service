package service

import (
	"errors"

	"github.com/izzanzahrial/library-service/entity"
)

var bookID int

type library struct {
	books map[string]*entity.BookQty
}

func (l *library) RegisterBook(title, author string, qty int) error {
	bookID += 1

	l.books[title] = &entity.BookQty{
		Book: entity.Book{
			ID:     bookID,
			Title:  title,
			Author: author,
		},
		Qty: qty,
	}

	return nil
}

func (l *library) GetBook(title string) (*entity.BookQty, error) {
	book, ok := l.books[title]
	if !ok {
		return &entity.BookQty{}, errors.New("cannot find the book")
	}

	if book.Qty <= 0 {
		return &entity.BookQty{}, errors.New("cannot borrow the book because the stock is empty")
	}

	book.Qty -= 1

	return book, nil
}

func (l *library) ReturnBook(title string) error {
	book, ok := l.books[title]
	if !ok {
		return errors.New("cannot return the book, because the book is not registerd")
	}

	book.Qty += 1

	return nil
}

func (l *library) ChangeBookQty(title string, qty int) error {
	book, err := l.GetBook(title)
	if err != nil {
		return err
	}

	if qty < 0 {
		book.Qty -= qty
		return nil
	}

	book.Qty += qty

	return nil
}
