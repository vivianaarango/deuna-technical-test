package main

func main() {
	// create a service struct.
	app := &Service{}

	// initialize api.
	app.Initialize()

	// run the server with specific port.
	app.Run(":8080")
}
