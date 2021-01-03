package models

import (
	"fmt"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type Cart struct {
	BookIds []int

	Books []Book
}

func (order Cart) Validate(v *revel.Validation) {
}

func (o Cart) String() string {
	return fmt.Sprintf("Cart(%v)", o.Books)
}

func (o *Cart) PreInsert(_ gorp.SqlExecutor) error {
	for i, book := range o.Books {
		o.BookIds[i] = book.BookId
	}
	return nil
}

func (o *Cart) PostGet(exe gorp.SqlExecutor) error {
	// var (
	// 	obj interface{}
	// 	err error
	// )

	// obj, err = exe.Get(User{}, o.UserId)
	// if err != nil {
	// 	return fmt.Errorf("Error loading a order's user (%d): %s", o.UserId, err)
	// }
	// o.User = obj.(*User)

	// obj, err = exe.Get(Book{}, o.BookId)
	// if err != nil {
	// 	return fmt.Errorf("Error loading a order's hotel (%d): %s", o.BookId, err)
	// }
	// o.Book = obj.(*Book)

	return nil
}
