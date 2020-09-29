package handlers

import (
	"ws/models"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	// _ "github.com/lib/pq"
)

func ProductsList(c *fiber.Ctx) error {
	product := models.Product{}
	// err := db.Get(&product, "select * from products where id=$1", "1")
	// if err != nil {
	// fmt.Printf("[error] %v\n", err)
	// }
	data := struct {
		Session models.Session
	}{
		Session: models.Session{
			UserName:          "Lucas",
			CartProductsCount: 3,
			Categories:        []string{"notebook", "monitor"},
		},
	}
	data.Session.UserGroups.Set(models.GroupAdmin)
	log.Printf("Is admin: %v", data.Session.UserGroups.IsAdmin())

	fmt.Printf("[debug] %+v", product)
	// return c.Render("admin/productList", fiber.Map{
	// "Title": "Hello, World!",
	// })
	return c.Render("admin/productList", data)
}
