package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/happynet78/go-fiber-blog/database"
	"github.com/happynet78/go-fiber-blog/model"
	"log"
	"time"
)

// Blog List
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog list",
	}

	time.Sleep(time.Millisecond * 5000)

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

// Blog detail page
func BlogDetail(c *fiber.Ctx) error {
	c.Status(200)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.Id == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		c.Status(404)
		return c.JSON(context)
	}

	context["record"] = record
	context["statusText"] = "Ok"
	context["msg"] = "Blog Detail"
	c.Status(200)
	return c.JSON(context)
}

// Add a blog into Database
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog list",
	}

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := database.DBConn.Create(&record)

	if result.Error != nil {
		log.Println("Error in saving data")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved successfully"
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// Update a Blog
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	// http://localhost:8000/blog/2

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.Id == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."

		c.Status(404)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := database.DBConn.Save(&record)

	if result.Error != nil {
		log.Println("Error in saving data")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is updated successfully"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

// Delete a Blog
func BlogDelete(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.Id == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Delete(&record)

	if result.Error != nil {
		context["msg"] = "Something went wrong."

		c.Status(400)
		return c.JSON(context)
	}

	context["statusText"] = "Ok"
	context["msg"] = "Record deleted successfully."
	c.Status(200)
	return c.JSON(context)
}
