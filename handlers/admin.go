package handlers

import (
	"fmt"
	"log"
	"ws/db"
	"ws/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // postgres drive
)

// Get product list
func ProductsList(c *fiber.Ctx) error {
	products, err := db.GetAllProduct()
	if err != nil {
		fmt.Printf("[error] %v\n", err)
	}
	data := struct {
		Session  models.Session
		Products []models.Product
	}{
		Session: models.Session{
			UserName:          "Lucas",
			CartProductsCount: 3,
			Categories:        []string{"notebook", "monitor"},
		},
		Products: products,
	}
	data.Session.UserGroups.Set(models.GroupAdmin)
	log.Printf("Is admin: %v", data.Session.UserGroups.IsAdmin())

	fmt.Printf("[debug] products: %+v", products)
	return c.Render("admin/productList", data)
}

// Get product item
func Product(c *fiber.Ctx) error {
	product, err := db.GetProductById("1")
	if err != nil {
		fmt.Printf("[error] %v\n", err)
	}
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
