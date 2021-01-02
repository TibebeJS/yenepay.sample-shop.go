package models

import (
	"fmt"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type Order struct {
	OrderId             int
	UserId              int
	BookId              int

	// Transient
	PurchaseDate   time.Time
	PublishingDate time.Time
	User           *User
	Book           *Book
}

func (order Order) Validate(v *revel.Validation) {
	v.Required(order.User)
	v.Required(order.Book)
	v.Required(order.PurchaseDate)
	v.Required(order.PublishingDate)
}

func (o Order) String() string {
	return fmt.Sprintf("Order(%s,%s)", o.User, o.Book)
}

func (o *Order) PreInsert(_ gorp.SqlExecutor) error {
	o.UserId = o.User.UserId
	o.BookId = o.Book.BookId
	return nil
}

func (o *Order) PostGet(exe gorp.SqlExecutor) error {
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

	return nil
}
