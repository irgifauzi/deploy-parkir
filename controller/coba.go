// File controller.go
package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	gile "github.com/irgifauzi/backend-parkir"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetTempatData(c *fiber.Ctx) error {
	ps := gile.GetAllTempat()
	return c.JSON(ps)
	
}