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
