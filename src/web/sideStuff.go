package web

import (
	"fmt"
	"strings"
	"website/src/config"

	"github.com/gofiber/fiber/v2"
	fiberLog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetLogLevel() {
	fiberLog.SetLevel(fiberLog.LevelInfo)
}

var (
	CFG    config.CFG = *config.GetConfig()                                      // The config
	server string     = fmt.Sprintf("%s:%d", CFG.Website.Host, CFG.Website.Port) // host:port

	APP *fiber.App = fiber.New(fiber.Config{
		ServerHeader: "technulgy:fiber",
		AppName:      "Technulgy",
	}) // Fiber App

	c = cors.New(cors.Config{
		AllowOrigins: strings.Join([]string{
			"https://technulgy.com",
			"https://*.technulgy.com",
			"http://localhost",
		}, ","),

		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
		}, ","),

		AllowHeaders: strings.Join([]string{
			"application/json",
		}, ","),

		AllowCredentials: true,
	}) // Cors Policy

	mon = monitor.New(monitor.Config{
		Title: "Technulgy Monitor",
	}) // Monitor

	auth = basicauth.New(basicauth.Config{
		Next: checkAuth,
		Users: map[string]string{
			"technulgy": "12345678",
		},
		Realm:           "Forbidden",
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}) // Basic Auth

	err error
)
