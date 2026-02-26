package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// The rest of your handlers go here...

// GetHabits - Fetches all habits from the database
func Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var data Request
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user := User{
		Email:    data.Email,
		Password: string(password),
	}

	DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	// For now, we'll keep it simple. Usually, you'd return a JWT token here.
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var data Request
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user User
	DB.Where("email = ?", data.Email).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Invalid Credentials"})
	}

	return c.JSON(user) // Return user info to the frontend
}
func GetHabits(c *fiber.Ctx) error {
	var habits []Habit

	// DB.Find picks up all columns, including the new 'Icon' field
	result := DB.Find(&habits)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch habits"})
	}

	return c.JSON(habits)
}

// CreateHabit - Adds a new habit with its name and icon
func CreateHabit(c *fiber.Ctx) error {
	habit := new(Habit)

	// BodyParser automatically maps JSON fields (name, icon) to the Habit struct
	if err := c.BodyParser(habit); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Save the new habit to PostgreSQL
	result := DB.Create(&habit)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save habit"})
	}

	return c.Status(201).JSON(habit)
}

// UpdateHabit - Mark as done/undone
func UpdateHabit(c *fiber.Ctx) error {
	id := c.Params("id")
	var habit Habit
	DB.First(&habit, id)
	habit.Completed = !habit.Completed // Switch between true/false
	DB.Save(&habit)
	return c.JSON(habit)
}

// DeleteHabit - Remove from DB
func DeleteHabit(c *fiber.Ctx) error {
	id := c.Params("id")
	DB.Delete(&Habit{}, id)
	return c.SendStatus(204)
}

// ResetHabits - Sets all habits to not completed
func ResetHabits(c *fiber.Ctx) error {
	// Updates all habits in the table to completed = false
	result := DB.Model(&Habit{}).Where("completed = ?", true).Update("completed", false)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not reset habits"})
	}
	return c.JSON(fiber.Map{"message": "Daily reset successful!"})
}
