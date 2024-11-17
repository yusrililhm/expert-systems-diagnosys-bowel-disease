package main

import "usus-sehat/internal/app"

func main() {
	go app.StartNonTLSServer()

	app.StartApp()
}
