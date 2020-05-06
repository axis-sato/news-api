package handler

import (
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func ErrorHandler(c *fiber.Ctx, err error) {
	log.WithFields(log.Fields{
		"error": err,
	}).Error("Internal Server Error")

	_ = c.JSON(newInternalServerErrorResponse())
	c.SendStatus(500)
}