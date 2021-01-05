package models

import (
	"fmt"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type CartItem struct {
	BookId int
	UserId int
	Quantity int

	Price float64

	Book *Book
	User *User
}

func (order CartItem) Validate(v *revel.Validation) {
}

func (o CartItem) String() string {
	return fmt.Sprintf("CartItem(%s, %s)", o.User.Name, o.Book.Title)
}

func (o *CartItem) PreInsert(_ gorp.SqlExecutor) error {
	o.BookId = o.Book.BookId
	o.UserId = o.User.UserId
	return nil
}

func (o *CartItem) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)

	obj, err = exe.Get(User{}, o.UserId)
	if err != nil {
		return fmt.Errorf("Error loading a order's user (%d): %s", o.UserId, err)
	}
	o.User = obj.(*User)

	obj, err = exe.Get(Book{}, o.BookId)
	if err != nil {
		return fmt.Errorf("Error loading a order's hotel (%d): %s", o.BookId, err)
	}
	o.Book = obj.(*Book)

	o.Price = o.Book.Price * float64(o.Quantity)

	return nil
}
