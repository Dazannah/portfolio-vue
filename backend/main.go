package main

func main() {
	router := registerRoutes()

	router.Run("localhost:8000")
}
