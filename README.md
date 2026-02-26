# Wilber's Habit-Tracker 

A modern, full-stack habit tracking application designed to help users build consistency through visual progress and smart automation.

##  Key Features
* **User Authentication:** Secure login and registration using Bcrypt password hashing.
* **Smart Icon Suggestion:** Automatically detects keywords (e.g., "water", "gym") to assign relevant emojis.
* **Progress Visualization:** Dynamic progress bar that updates in real-time as habits are completed.
* **Responsive UI:** A clean, card-based dashboard optimized for productivity.

##  Tech Stack
* **Frontend:** Vue 3 (Composition API), Vite, Axios.
* **Backend:** Go (Golang), Fiber Web Framework.
* **Database:** PostgreSQL with GORM (Object-Relational Mapper).
* **Security:** JWT-ready structure with hashed password storage.

## How to Run Locally
1. **Database:** Ensure PostgreSQL is running and create a database named `habits`.
2. **Backend:** ```bash
   cd backend
   go run .
