package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	connStr := "postgresql://postgres:@localhost/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("/", "./public")

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		log.Fatalln(err)
		c.JSON("An error has ocurred")
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString(("Hello: post"))
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString(("Hello: put"))
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString(("Hello: delete"))
}
