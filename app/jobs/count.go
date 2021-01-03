package jobs

import (
	"fmt"

	"github.com/revel/revel"

	"github.com/TibebeJs/yenepay.sample-shop.go/app/models"
	"github.com/revel/modules/jobs/app/jobs"
	gorp "github.com/revel/modules/orm/gorp/app"
)

// Periodically count the orders in the database.
type OrderCounter struct{}

func (c OrderCounter) Run() {
	orders, err := gorp.Db.Map.Select(models.Order{},
		`select * from "Order"`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d orders.\n", len(orders))
}

func init() {
	revel.OnAppStart(func() {
		err := jobs.Schedule("@every 6s", OrderCounter{})
		if err != nil {
			fmt.Print(err)
		}
	})
}
