package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jmoiron/sqlx"
)

type RunMode int

const (
	Dev RunMode = iota
	Prod
	Test
)

const (
	version string = "0.1.0"
)

func (r RunMode) String() string {
	return [...]string{"development", "production", "test"}[r]
}

var mode = Prod

var db *sqlx.DB

func init() {
	// Run mode.
	mode = Dev
	if strings.HasPrefix(strings.ToLower(os.Getenv("RUN_MODE")), "prod") {
		mode = Prod
	}
	log.Printf("Running in %v mode (version %s)\n", mode, version)

}

func main() {

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=ws dbname=ws sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	engine := html.New("./views", ".html")
	if mode == Dev {
		engine.Reload(true) // Optional. Default: false
		engine.Debug(true)  // Optional. Default: false
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		fmt.Println("ping")
		return c.SendString("pong")
	})

	app.Get("/product", func(c *fiber.Ctx) error {
		product := Product{}
		err := db.Get(&product, "select * from products where id=$1", "1")
		if err != nil {
			fmt.Printf("[error] %v\n", err)
		}
		fmt.Printf("[debug] %+v", product)
		return c.SendString(product.Title)
	})

	app.Get("/admin/products", func(c *fiber.Ctx) error {
		product := Product{}
		err := db.Get(&product, "select * from products where id=$1", "1")
		if err != nil {
			fmt.Printf("[error] %v\n", err)
		}
		data := struct {
			Session Session
		}{
			Session: Session{
				UserName:          "Lucas",
				CartProductsCount: 3,
			},
		}
		data.Session.UserGroups.Set(GroupAdmin)
		log.Printf("Is admin: %v", data.Session.UserGroups.IsAdmin())

		fmt.Printf("[debug] %+v", product)
		// return c.Render("admin/productList", fiber.Map{
		// "Title": "Hello, World!",
		// })
		return c.Render("admin/productList", data)
	})

	app.Static("/", "./static")

	log.Fatal(app.Listen(":8080"))
}
