package auth

import (
	"fmt"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Cache state
var app *fiber.App

// init
func init() {

	// Initialize app
	app = fiber.New(fiber.Config{
		AppName: "Test App v1.0.1",
	})
	
	app.Use(logger.New(
		logger.Config{
			Format: "[${time}] ${status} - ${latency} ${method} ${path} - ${header:}\n​",
		},
	))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Access-Control-Allow-Headers, X-Requested-With, X-HTTP-Method-Override, access-control-allow-origin, access-control-allow-headers",
	}))

	app.Get("/login", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World from auth.login 👋 !")
    })
	app.Get("/login/sign/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
        return c.SendString(fmt.Sprintf("Hello, World from auth.login.sign 👋 ! %s",name))
    })
}

	
// This is a simple REST API that responds with a personalized greeting.
// To call it, run in your terminal:
//
//     curl http://localhost:4000/auth/login
//     curl http://localhost:4000/auth/login/sign/name-param
//
//encore:api public raw path=/auth/*p1
func Handle(w http.ResponseWriter, r *http.Request) {
	RemoveBaseURLFromRequest(r)
	adaptor.FiberApp(app)(w, r)
}
