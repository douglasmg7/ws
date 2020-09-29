package main

import (
	"ws/handlers"

	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

func main() {
	// Run mode.
	mode = Dev
	if strings.HasPrefix(strings.ToLower(os.Getenv("RUN_MODE")), "prod") {
		mode = Prod
	}
	log.Printf("Running in %v mode (version %s)\n", mode, version)

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

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	index := app.Group("/")
	admin := app.Group("/admin")

	// Index
	index.Get("/", handlers.Index)
	index.Get("/ping", handlers.Ping)

	// Admin
	admin.Get("/productslist", handlers.ProductsList)

	// Statics
	app.Static("/", "./static")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port
	log.Fatal(app.Listen(":8080"))
}
