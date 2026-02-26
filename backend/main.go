package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Connect to DB (This now migrates both Users and Habits)
	ConnectDB()

	// 2. Initialize Fiber App
	app := fiber.New()

	// 3. Enable CORS
	// Important: This must come BEFORE your routes
	app.Use(cors.New())

	// 4. Auth Routes
	app.Post("/api/register", Register)
	app.Post("/api/login", Login)

	// 5. Habit Routes
	app.Get("/api/habits", GetHabits)          // Gets habits filtered by userId
	app.Post("/api/habits", CreateHabit)       // Add a new habit
	app.Put("/api/habits/:id", UpdateHabit)    // Toggle complete
	app.Delete("/api/habits/:id", DeleteHabit) // Remove habit
	app.Post("/api/habits/reset", ResetHabits) // Daily reset

	// 6. Start Server
	log.Fatal(app.Listen(":3000"))
}
