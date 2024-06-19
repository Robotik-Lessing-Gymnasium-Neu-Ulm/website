package web

import (
	"log"
)

func StartTheWeb() {
	APP.Use(c)    // Cors middleware
	APP.Use(auth) // Basic auth for monitor side

	APP.Static("/", "./public") // Serve side

	APP.Get("/monitor", mon)

	err = APP.Listen(server) // Start server
	if err != nil {
		log.SetFlags(log.LstdFlags & log.Lshortfile)
		log.Printf("Error starting WebServer: %v\n", err)
	}
}
