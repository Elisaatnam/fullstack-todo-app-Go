package main

import (
	"fmt"
	"log"

	"slices"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("Hello, world!!!")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(200).JSON(todo)
	})


	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// todos[i].Completed = !todos[i].Completed
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// // todos[:i] creates a new slice containing all elements from the beginning of todos up to (but not including) index i
				// // todos[i+1:] creates another slice containing all elements from index i+1 to the end of todos
				// // append(todos[:i], todos[i+1:]...) combines these two slices, effectively skipping the element at index i
				// // The ... (ellipsis) operator unpacks the second slice so its elements can be individually appended
				// //The result is assigned back to the todos variable, replacing the original slice
				todos = slices.Delete(todos, i, i+1)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":4000"))
}

// "&" => memory adress
// "*" => pointer (to the value)