package controllers

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/revel/revel"

	"github.com/TibebeJs/yenepay.sample-shop.go/app/models"

	"github.com/Masterminds/squirrel"
)

type Books struct {
	Application
}

func (c Books) checkUser() revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}

	return nil
}

func (c Books) Index() revel.Result {
	c.Log.Info("Fetching books")
	var books []*models.Book
	_, err := c.Txn.Select(&books,
		c.Db.SqlStatementBuilder.Select("*").From("\"Book\""),
	)

	if err != nil {
		panic(err)
	}

	return c.Render(books)
}

func (c Books) List(search string, size, page uint64) revel.Result {
	if page == 0 {
		page = 1
	}
	nextPage := page + 1
	search = strings.TrimSpace(search)

	var books []*models.Book
	builder := c.Db.SqlStatementBuilder.Select("*").From("\"Book\"").Offset((page - 1) * size).Limit(size)
	if search != "" {
		search = "%" + strings.ToLower(search) + "%"
		builder = builder.Where(squirrel.Or{
			squirrel.Expr("lower(Name) like ?", search),
			squirrel.Expr("lower(City) like ?", search)})
	}
	if _, err := c.Txn.Select(&books, builder); err != nil {
		c.Log.Fatal("Unexpected error loading books", "error", err)
	}

	return c.Render(books, search, size, page, nextPage)
}

func (c Books) loadBookById(id int) *models.Book {
	h, err := c.Txn.Get(models.Book{}, id)
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.Book)
}

func (c Books) Show(id int) revel.Result {
	book := c.loadBookById(id)
	if book == nil {
		return c.NotFound("Book %d does not exist", id)
	}
	title := book.Title
	return c.Render(title, book)
}

func (c Books) Settings() revel.Result {
	return c.Render()
}

func (c Books) SaveSettings(password, verifyPassword string) revel.Result {
	models.ValidatePassword(c.Validation, password)
	c.Validation.Required(verifyPassword).
		Message("Please verify your password")
	c.Validation.Required(verifyPassword == password).
		Message("Your password doesn't match")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		return c.Redirect(Books.Settings)
	}

	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := c.Txn.ExecUpdate(c.Db.SqlStatementBuilder.
		Update("User").Set("HashedPassword", bcryptPassword).
		Where("UserId=?", c.connected().UserId))
	if err != nil {
		panic(err)
	}
	c.Flash.Success("Password updated")
	return c.Redirect(Books.Index)
}

func (c Books) ConfirmOrder(id int, order models.Order) revel.Result {
	book := c.loadBookById(id)
	if book == nil {
		return c.NotFound("Book %d does not exist", id)
	}

	title := fmt.Sprintf("Confirm %s order", book.Title)
	order.Book = book
	order.User = c.connected()
	order.Validate(c.Validation)

	if c.Validation.HasErrors() || c.Params.Get("revise") != "" {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Books.Book, id)
	}

	if c.Params.Get("confirm") != "" {
		err := c.Txn.Insert(&order)
		if err != nil {
			panic(err)
		}
		c.Flash.Success("Thank you, %s, your confirmation number for %s is %d",
			order.User.Name, book.Title, order.OrderId)
		return c.Redirect(Books.Index, id)
	}

	return c.Render(title, book, order)
}

func (c Books) CancelOrder(id int) revel.Result {
	_, err := c.Txn.Delete(&models.Order{OrderId: id})
	if err != nil {
		panic(err)
	}
	c.Flash.Success(fmt.Sprintln("Order cancelled for confirmation number", id))
	return c.Redirect(Books.Index)
}

func (c Books) Book(id int) revel.Result {
	book := c.loadBookById(id)
	if book == nil {
		return c.NotFound("Book %d does not exist", id)
	}

	title := "Book " + book.Title
	return c.Render(title, book)
}
