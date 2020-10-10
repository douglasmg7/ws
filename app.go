package main

import (
	"ws/db"
	"ws/handlers"
	"ws/models"

	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/session/v2/provider/redis"
	"github.com/gofiber/template/html"
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

var sessions *session.Session

func main() {
	// Run mode.
	mode = Dev
	if strings.HasPrefix(strings.ToLower(os.Getenv("RUN_MODE")), "prod") {
		mode = Prod
	}
	log.Printf("Running in %v mode (version %s)\n", mode, version)

	db.Connect()
	defer db.Close()

	app := SetupApp()

	// Listen on port
	log.Fatal(app.Listen(":8080"))
}

func SetupApp() *fiber.App {
	// Engine template
	engine := html.New("./views", ".html")
	if mode == Dev {
		engine.Reload(true) // Optional. Default: false
		engine.Debug(true)  // Optional. Default: false
	}

	// Error handler
	errorHandler := handlers.ErrorHandler
	if mode == Prod {
		errorHandler = handlers.ErrorHandlerProd
	}

	// Fiber configuration
	app := fiber.New(fiber.Config{
		Prefork:      false,
		Views:        engine,
		ErrorHandler: errorHandler,
	})

	// Session
	provider, err := redis.New(redis.Config{
		KeyPrefix:   "session",
		Addr:        "127.0.0.1:6379",
		PoolSize:    8,
		IdleTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	sessions = session.New(session.Config{
		Provider: provider,
	})

	///////////////////////////////////////////////////////////////////////////
	// MIDDLEWARE
	///////////////////////////////////////////////////////////////////////////
	// Recovery
	app.Use(recover.New())
	// Logger
	app.Use(logger.New())
	// Limiter
	app.Use(limiter.New(limiter.Config{
		Duration: 5 * time.Second,
		Max:      30,
	}))
	// Favicon
	app.Use(favicon.New(favicon.Config{
		File: "./static/logo.png",
	}))
	// Session
	app.Use(sessionMiddleware)

	index := app.Group("/")
	admin := app.Group("/admin")

	///////////////////////////////////////////////////////////////////////////
	// ROUTER
	///////////////////////////////////////////////////////////////////////////
	// Index
	index.Get("/", handlers.Index)
	index.Get("/ping", handlers.Ping)

	// Admin
	admin.Get("/pid", handlers.Pid)
	admin.Get("/session", handlers.Session)
	admin.Get("/products", handlers.ProductsList)

	// Statics
	app.Static("/", "./static")

	// Handle not founds
	app.Use(handlers.NotFound)

	return app
}

func sessionMiddleware(c *fiber.Ctx) error {
	store := sessions.Get(c) // get/create new session
	defer store.Save()

	// log.Printf("Session id: %v, User id: %v", store.ID(), store.Get("user_id"))

	// store.ID()               // returns session id
	// store.Destroy() // delete storage + cookie
	// store.Get("john")        // get from storage
	// store.Regenerate() // generate new session id
	// store.Delete("john")     // delete from storage
	store.Set("user_id", "Fernando") // save to storage

	// return c.SendString(fmt.Sprintf("Session id: %v, User id: %v", store.ID(), store.Get("user_id")))

	c.Locals("session_id", store.ID())

	session := models.Session{
		UserName:          "Lucas",
		CartProductsCount: 3,
		Categories:        []string{"notebook", "monitor"},
	}
	session.UserGroups.Set(models.GroupAdmin)

	c.Locals("session", session)

	return c.Next()
}
