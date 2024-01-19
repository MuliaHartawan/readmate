package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Routing dan middleware dapat ditambahkan di sini

    app.Listen(":3000")
}

