package service

import (
	"fmt"

	"github.com/izzanzahrial/library-service/entity"
)

type borrowing struct {
	user      User
	library   Library
	borrowing map[int]*entity.Borrow
}

func (b *borrowing) BorrowBooks(titles []string, user *entity.User) error {
	borrowKey := userID

	b.borrowing[borrowKey] = &entity.Borrow{
		User: *user,
	}

	for _, title := range titles {
		book, _ := b.library.GetBook(title)
		b.borrowing[borrowKey] = &entity.Borrow{
			Books: *book,
		}
	}

	return nil
}

func (b *borrowing) ReturnBooks(books []*entity.Book, user *entity.User) error {
	if _, ok := b.borrowing[user.ID]; !ok {
		return fmt.Errorf("user with the name %s havent borrow a book", user.Name)
	}

	for _, book := range books {
		err := b.library.ReturnBook(book.Title)
		if err != nil {
			return err
		}
	}

	delete(b.borrowing, user.ID)

	return nil
}
