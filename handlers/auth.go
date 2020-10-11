package handlers

import (
	// "fmt"
	// "log"
	// "ws/db"
	"ws/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // postgres drive
)

type signinData struct {
	Session          models.Session
	Email            valueMessage
	Password         valueMessage
	WarnMsgHead      string
	SuccessMsgHead   string
	WarnMsgFooter    string
	SuccessMsgFooter string
}

///////////////////////////////////////////////////////////////////////////////
// SIGNIN
///////////////////////////////////////////////////////////////////////////////

// Signin page
func Signin(c *fiber.Ctx) error {
	return c.Render("auth/signin", signinData{
		Session: c.Locals("session").(models.Session),
	})
}

// // Get product item
// func Signout(c *fiber.Ctx) error {
// product, err := db.GetProductById("1")
// if err != nil {
// fmt.Printf("[error] %v\n", err)
// }
// data := struct {
// Session models.Session
// }{
// Session: models.Session{
// UserName:          "Lucas",
// CartProductsCount: 3,
// Categories:        []string{"notebook", "monitor"},
// },
// }
// data.Session.UserGroups.Set(models.GroupAdmin)
// log.Printf("Is admin: %v", data.Session.UserGroups.IsAdmin())

// fmt.Printf("[debug] %+v", product)
// // return c.Render("admin/productList", fiber.Map{
// // "Title": "Hello, World!",
// // })
// return c.Render("admin/productList", data)
// }
