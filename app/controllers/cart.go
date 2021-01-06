package controllers

import (
	"fmt"
	"strconv"

	"github.com/TibebeJs/yenepay.sample-shop.go/app/models"
	"github.com/TibebeJs/yenepay.sdk.go/checkout"
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
		c.Flash.Error(fmt.Sprintf("Error with quantity: %s", err))
		return c.Redirect(Application.Index)
	} else if quantity < 1 {
		c.Flash.Error("Error with quantity: must be atleast 1")
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

func (c Cart) ExpressCheckout() revel.Result {
	var (
		err      error
		bookId   int
		quantity int
	)

	yenepay := checkout.NewYenePayCheckOut()

	bookId, err = strconv.Atoi(c.Params.Form.Get("book_id"))
	if err != nil {
		return c.Redirect(Application.Index)
	}

	quantity, err = strconv.Atoi(c.Params.Form.Get("quantity"))

	if err != nil {
		c.Flash.Error(fmt.Sprintf("Error with quantity: %s", err))
		return c.Redirect(Application.Index)
	} else if quantity < 1 {
		c.Flash.Error("Error with quantity: must be atleast 1")
		return c.Redirect(Application.Index)
	}

	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}

	h, err := c.Txn.Get(models.Book{}, bookId)
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}

	var book *models.Book = h.(*models.Book)

	url := yenepay.ExpressCheckoutUrl(checkout.NewCheckoutOption(
		checkout.OptionsParams{
			UseSandbox:      true,
			Process:         checkout.ExpressCheckout,
			MerchantId:      "0694",
			SuccessUrl:      "http://localhost:3000/payments/PaymentSuccessReturnUrl",
			CancelUrl:       "http://localhost:3000/payments/CancelReturnUrl",
			IPNUrl:          "http://localhost:3000/payments/IPNUrl",
			FailureUrl:      "http://localhost:3000/payments/FailureUrl",
			ExpiresAfter:    2880,
			MerchantOrderId: "ab-cd",
		},
	), checkout.NewExpressCheckoutItem(
		checkout.ExpressParams{ItemId: fmt.Sprint(bookId), ItemName: book.Title, UnitPrice: book.Price, Quantity: quantity},
	))

	return c.Redirect(url)
}
