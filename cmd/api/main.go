package main

import "simple-bank/internal/routes"

func main() {
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
