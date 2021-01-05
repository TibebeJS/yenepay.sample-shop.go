package controllers

import (
	"strconv"

	"github.com/TibebeJs/yenepay.sample-shop.go/app/models"
	"github.com/revel/revel"
)

type Cart struct {
	Application
}

func (c Cart) RemoveItem() revel.Result {
	bookId, err := strconv.Atoi(c.Params.Query.Get("book_id"))

	if err != nil {
		return c.Redirect(Application.Index)
	}

	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}

	_, err = c.Txn.Delete(&models.CartItem{UserId: user.UserId, BookId: bookId})
	if err != nil {
		c.Flash.Error("Couldn't remove a cart item")
		c.Log.Fatal("Unexpected error removing a cart item", "error", err)
	}

	return c.Redirect(Application.Index)
}

func (c Cart) AddItem() revel.Result {
	var (
		err      error
		bookId   int
		quantity int
	)

	bookId, err = strconv.Atoi(c.Params.Form.Get("book_id"))
	if err != nil {
		return c.Redirect(Application.Index)
	}

	quantity, err = strconv.Atoi(c.Params.Form.Get("quantity"))

	if err != nil {
		return c.Redirect(Application.Index)
	}

	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}

	err = c.Db.Map.Insert(&models.CartItem{Quantity: quantity, User: &models.User{UserId: user.UserId}, Book: &models.Book{BookId: bookId}})

	if err != nil {
		c.Flash.Error("Couldn't insert a cart item")
		c.Log.Fatal("Unexpected error inserting a cart item", "error", err)
	}

	return c.Redirect(Application.Index)
}
