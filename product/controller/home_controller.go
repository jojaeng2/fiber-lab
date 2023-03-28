package controller

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

type HomeController interface {
	GetHome(c *fiber.Ctx) error
}

type HomeControllerImpl struct {
}

func NewHomeController() HomeController {
	return &HomeControllerImpl{}
}

func (controller *HomeControllerImpl) GetHome(c *fiber.Ctx) error {
	tpl, err := template.ParseFiles("views/home.html")
	if err != nil {
		return err
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Homepage",
		Message: "Hello, World!!!!",
	}

	var bodyContent bytes.Buffer
	if err := tpl.Execute(&bodyContent, data); err != nil {
		return err
	}
	c.Set("Content-Type", "text/html")
	return c.SendString(bodyContent.String())
}
