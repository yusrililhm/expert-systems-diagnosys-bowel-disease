package main

import "usus-sehat/app"

func main() {
	go app.StartNonTLSServer()
	
	app.StartApp()
}
