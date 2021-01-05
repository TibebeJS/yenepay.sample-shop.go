package app

import (
	"time"

	"github.com/TibebeJs/yenepay.sample-shop.go/app/models"
	"github.com/go-gorp/gorp"
	rgorp "github.com/revel/modules/orm/gorp/app"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(func() {
		Dbm := rgorp.Db.Map
		setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
			for col, size := range colSizes {
				t.ColMap(col).MaxSize = size
			}
		}

		t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
		t.ColMap("Password").Transient = true
		setColumnSizes(t, map[string]int{
			"Username": 20,
			"Name":     100,
		})

		t = Dbm.AddTable(models.CartItem{}).SetKeys(false, "UserId", "BookId")
		t.ColMap("User").Transient = true
		t.ColMap("Book").Transient = true
		t.ColMap("Price").Transient = true

		t = Dbm.AddTable(models.Book{}).SetKeys(true, "BookId")
		setColumnSizes(t, map[string]int{
			"Title":      50,
			"Price":      100,
			"CoverImage": 60,
		})

		t = Dbm.AddTable(models.Order{}).SetKeys(true, "OrderId")
		t.ColMap("User").Transient = true
		t.ColMap("Book").Transient = true

		rgorp.Db.TraceOn(revel.AppLog)
		err := Dbm.DropTablesIfExists()

		if err != nil {
			revel.AppLog.Fatal("Failed to drop existing tables", "error", err)
		}

		err = Dbm.CreateTablesIfNotExists()
		if err != nil {
			revel.AppLog.Fatal("Failed to create tables", "error", err)
		}

		demoUsername := "user"
		demoPassword := "password"

		bcryptPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(demoPassword), bcrypt.DefaultCost)

		demoUser := &models.User{
			UserId:         0,
			Name:           "Abebe Kebede",
			Username:       demoUsername,
			Password:       demoUsername,
			HashedPassword: bcryptPassword,
		}
		if err := Dbm.Insert(demoUser); err != nil {
			panic(err)
		}
		count, _ := rgorp.Db.SelectInt(rgorp.Db.SqlStatementBuilder.Select("count(*)").From("\"User\""))
		if count > 1 {
			revel.AppLog.Panic("Unexpected multiple users", "count", count)
		}

		books := []*models.Book{
			{
				BookId:     0,
				Price:      19.9,
				Title:      "Go Bootcamp: 2nd Edition",
				CoverImage: "go-bootcamp.jpg",
			},
			{
				BookId:     0,
				Price:      19.9,
				Title:      "Automate The Boring Stuff With Python",
				CoverImage: "automate-the-boring-stuff-with-python.jpg",
			},
			{
				BookId:     0,
				Price:      19.9,
				Title:      "Get Programming With Go",
				CoverImage: "get-programming-with-go.jpg",
			},
		}
		for _, book := range books {
			if err := Dbm.Insert(book); err != nil {
				panic(err)
			}
		}
		orders := []*models.Order{
			{
				OrderId:        0,
				UserId:         demoUser.UserId,
				BookId:         books[0].BookId,
				PurchaseDate:   time.Now(),
				PublishingDate: time.Now(),
				User:           demoUser,
				Book:           books[0],
			},
			{
				OrderId:        0,
				UserId:         demoUser.UserId,
				BookId:         books[1].BookId,
				PurchaseDate:   time.Now(),
				PublishingDate: time.Now(),
				User:           demoUser,
				Book:           books[1],
			},
			{
				OrderId:        0,
				UserId:         demoUser.UserId,
				BookId:         books[2].BookId,
				PurchaseDate:   time.Now(),
				PublishingDate: time.Now(),
				User:           demoUser,
				Book:           books[2],
			},
		}
		for _, order := range orders {
			if err := Dbm.Insert(order); err != nil {
				panic(err)
			}
		}

		cartItem := models.CartItem{
			UserId: demoUser.UserId,
			BookId: 0,
			Book:   books[0],
			User:   demoUser,
			Quantity: 2,
		}

		if err := Dbm.Insert(&cartItem); err != nil {
			panic(err)
		}
	}, 5)

}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
