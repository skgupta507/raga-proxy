package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"codeberg.org/aryak/paattu-proxy/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:                 "paattu-proxy",
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0/0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	app.Get("/media/:id/:path", func(c *fiber.Ctx) error {
		utils.ProxyRequest(c, "https://c.saavncdn.com/"+c.Params("id")+"/"+c.Params("path"))
		return nil
	})
	app.Get("/aac/:id/:path", func(c *fiber.Ctx) error {
		utils.ProxyRequest(c, "https://c.saavncdn.com/"+c.Params("id")+"/"+c.Params("path"))
		return nil
	})
	app.Listen(GetPort())
}

// GetPort returns the port to listen on
func GetPort() string {
	port := os.Getenv("PAATTU_PROXY_PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}
