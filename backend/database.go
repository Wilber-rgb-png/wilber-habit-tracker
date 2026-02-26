package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Updated Habit Model with an Icon field
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"` // "-" means never send the password in JSON
}

type Habit struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Completed bool   `json:"completed" gorm:"default:false"`
	UserID    uint   `json:"user_id"` // Links the habit to a user
}

func ConnectDB() {
	// Connection string - ensure YOUR_PASSWORD is correct
	dsn := "host=localhost user=postgres password=#clozet662004 dbname=habit_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")

	// 1. Sync the schema (Adds the Icon column automatically)
	db.AutoMigrate(&User{}, &Habit{})
	DB = db

	// 2. Seed the database if it's empty
	SeedHabits()
}

// SeedHabits adds the initial habits you asked for
func SeedHabits() {
	var count int64
	DB.Model(&Habit{}).Count(&count)

	if count == 0 {
		fmt.Println("🌱 Seeding initial habits...")
		initialHabits := []Habit{
			{Name: "Drink 2L Water", Icon: "💧", Completed: false},
			{Name: "30-Minute Walk", Icon: "🏃", Completed: false},
			{Name: "10-Minute Meditation", Icon: "🧘", Completed: false},
			{Name: "Read 10 Pages", Icon: "📚", Completed: false},
		}

		for _, h := range initialHabits {
			DB.Create(&h)
		}
		fmt.Println("✅ Seeding complete!")
	}
}
