package main

import "website/src/web"

func main() {
	web.SetLogLevel() // Sets the fiber log level
	web.StartTheWeb() // Starts the web site and all the api routes ...
}
