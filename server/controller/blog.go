package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/happynet78/go-fiber-blog/database"
	"github.com/happynet78/go-fiber-blog/model"
	"github.com/happynet78/go-fiber-blog/util"
	"log"
	"math/rand"
	"os"
	"time"
)

// Blog List
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog list",
	}

	time.Sleep(time.Millisecond * 1500)

	db := database.DBConn

	var records []model.Blog

	db.Find(&records).Order("id desc")

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

	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	fmt.Println("uid: " + id)
	// var blog model.Blog

	// blogRecord := new(model.Blog)

	// record := database.DBConn.Model(&blog).Where("user_id = ?", uid).Preload("User").Find(&blog)
	record := new(model.Blog)

	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	// File upload
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	fileName := ""

	for _, file := range files {

		fileName = randLetter(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./static/uploads/"+fileName); err != nil {
			return nil
		}

		// record.image = "../static/uploads/" + fileName
	}

	// record.userid = id
	log.Println(record)

	result := database.DBConn.Create(record)

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

	// Remove image
	filename := record.Image

	err := os.Remove(filename)
	if err != nil {
		log.Println("Error in deleting file.", err)
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

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
