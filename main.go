package main

import (
	"database/sql"
	"employee/config"
	"employee/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

var db *sql.DB
var err error

func init() {
	config.InitDB()
	db = config.DB

}

func main() {

	app := fiber.New()

	defer config.DB.Close()

	app.Post("/departments", addDepartment)
	app.Get("/departments", getAllDepartment)

	// Listen on a specific port (3000 in this case) and keep the server running
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

// get all departs func
func getAllDepartment(c *fiber.Ctx) error {

	// Declare and initialize the rows variable
	rows, err := db.Query(" select id, name from department order by name ")
	if err != nil {
		log.Fatal(err)
		return c.SendStatus(500)
	}
	defer rows.Close()

	var departments []model.Department

	for rows.Next() {
		var depart model.Department
		if err := rows.Scan(&depart.ID, &depart.Name); err != nil {
			log.Fatal(err)
			return c.SendStatus(500)
		}
		departments = append(departments, depart)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return c.SendStatus(500)
	}

	return c.JSON(departments)
}

// add department dunc

func addDepartment(c *fiber.Ctx) error {

	bodyString := string(c.Body())
	println("body", bodyString)

	depart := new(model.Department)
	if err := c.BodyParser(depart); err != nil {
		println("BodyParser: ", err)
		return c.SendStatus(500)
	}

	_, err = db.Exec("insert into department ( name) values (?) ", depart.Name)
	if err != nil {
		log.Fatal(err)
		return c.SendStatus(500)
	}

	return c.JSON(depart)

}
