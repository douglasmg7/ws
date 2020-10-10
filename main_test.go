package main

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func TestMain(m *testing.M) {
	setupTest()
	code := m.Run()
	shutdownTest()

	os.Exit(code)
}

func setupTest() {
	// initRedis()
	// initSql3DB()
	app = SetupApp()
}

func shutdownTest() {
	// closeRedis()
	// closeSql3DB()
}
