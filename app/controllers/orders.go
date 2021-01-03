package controllers

import (
	"github.com/revel/revel"

	"github.com/TibebeJs/yenepay.sample-shop.go/app/models"
)

type Orders struct {
	Application
}

func (c Orders) Index() revel.Result {
	c.Log.Info("Fetching orders")
	var orders []*models.Order
	if user := c.connected(); user != nil {
		builder := c.Db.SqlStatementBuilder.Select("*").From("\"Order\"").Where("\"UserId\"=?", user.UserId)
		if _, err := c.Txn.Select(&orders, builder); err != nil {
			c.Log.Fatal("Unexpected error loading orders", "error", err)
		}
	} else {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}

	return c.Render(orders)
}

// func (c Orders) loadOrderById(id int) *models.Order {
// 	h, err := c.Txn.Get(models.Order{}, id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if h == nil {
// 		return nil
// 	}
// 	return h.(*models.Order)
// }

// func (c Orders) Show(id int) revel.Result {
// 	order := c.loadOrderById(id)
// 	if order == nil {
// 		return c.NotFound("Order %d does not exist", id)
// 	}
// 	title := order.Title
// 	return c.Render(title, order)
// }

// func (c Orders) Settings() revel.Result {
// 	return c.Render()
// }

// func (c Orders) SaveSettings(password, verifyPassword string) revel.Result {
// 	models.ValidatePassword(c.Validation, password)
// 	c.Validation.Required(verifyPassword).
// 		Message("Please verify your password")
// 	c.Validation.Required(verifyPassword == password).
// 		Message("Your password doesn't match")
// 	if c.Validation.HasErrors() {
// 		c.Validation.Keep()
// 		return c.Redirect(Orders.Settings)
// 	}

// 	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	_, err := c.Txn.ExecUpdate(c.Db.SqlStatementBuilder.
// 		Update("User").Set("HashedPassword", bcryptPassword).
// 		Where("UserId=?", c.connected().UserId))
// 	if err != nil {
// 		panic(err)
// 	}
// 	c.Flash.Success("Password updated")
// 	return c.Redirect(Orders.Index)
// }

// func (c Orders) ConfirmOrder(id int, order models.Order) revel.Result {
// 	order := c.loadOrderById(id)
// 	if order == nil {
// 		return c.NotFound("Order %d does not exist", id)
// 	}

// 	title := fmt.Sprintf("Confirm %s order", order.Title)
// 	order.Order = order
// 	order.User = c.connected()
// 	order.Validate(c.Validation)

// 	if c.Validation.HasErrors() || c.Params.Get("revise") != "" {
// 		c.Validation.Keep()
// 		c.FlashParams()
// 		return c.Redirect(Orders.Order, id)
// 	}

// 	if c.Params.Get("confirm") != "" {
// 		err := c.Txn.Insert(&order)
// 		if err != nil {
// 			panic(err)
// 		}
// 		c.Flash.Success("Thank you, %s, your confirmation number for %s is %d",
// 			order.User.Name, order.Title, order.OrderId)
// 		return c.Redirect(Orders.Index, id)
// 	}

// 	return c.Render(title, order, order)
// }

// func (c Orders) CancelOrder(id int) revel.Result {
// 	_, err := c.Txn.Delete(&models.Order{OrderId: id})
// 	if err != nil {
// 		panic(err)
// 	}
// 	c.Flash.Success(fmt.Sprintln("Order cancelled for confirmation number", id))
// 	return c.Redirect(Orders.Index)
// }

// func (c Orders) Order(id int) revel.Result {
// 	order := c.loadOrderById(id)
// 	if order == nil {
// 		return c.NotFound("Order %d does not exist", id)
// 	}

// 	title := "Order " + order.Title
// 	return c.Render(title, order)
// }
