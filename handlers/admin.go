package handlers

import (
	"fmt"
	"log"
	"os"
	"ws/db"
	"ws/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // postgres drive
)

// Process ID
func Pid(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Worker #%v", os.Getpid()))
}

// Session
func Session(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Session id: %v", c.Locals("session_id")))
}

// Get product list
func ProductsList(c *fiber.Ctx) error {
	// return c.SendString("admin/productList")

	// Get products
	products, err := db.GetAllProduct()
	if err != nil {
		fmt.Printf("[error] %v\n", err)
	}

	// Data
	data := struct {
		Session  models.Session
		Products []models.Product
	}{
		Session:  c.Locals("session").(models.Session),
		Products: products,
	}

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
			CartProductsCount: 5,
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
