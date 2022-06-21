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

func (b *borrowing) BorrowBooks(titles []string, userID int) error {
	registeredUser, err := b.user.GetUser(userID)
	if err != nil {
		return err
	}

	var books []entity.Book

	for _, title := range titles {
		bookQty, _ := b.library.GetBook(title)
		books = append(books, bookQty.Book)
	}

	b.borrowing[registeredUser.ID] = &entity.Borrow{
		User:  *registeredUser,
		Books: books,
	}

	return nil
}

func (b *borrowing) ReturnBooks(titles []string, userID int) error {
	registeredUser, err := b.user.GetUser(userID)
	if err != nil {
		return err
	}

	for _, title := range titles {
		err := b.library.ReturnBook(title)
		if err != nil {
			fmt.Println(err)
		}
	}

	delete(b.borrowing, registeredUser.ID)

	return nil
}
