package webhook

import (
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/ws"
)

func PostMessage(c *fiber.Ctx) error {
	var h Header
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return err
	}
	var resp atmessage.Response
	var msg model.IteungMessage
	err = c.BodyParser(&msg)
	if err != nil {
		return err
	}

	// Handle login request
	if msg.Message == "Babi" || msg.Message == "Anjing" || msg.Message == "goblok" {
		resp = getRandomRude(msg)
	} else if msg.Message == "sayang" || msg.Message == "syg" || msg.Message == "cinta" {
		resp = getRandomRude(msg)
	} else if msg.Message == "cantik" {
		resp = getRandomCompliment(msg)
	} else if msg.Message == "Alice" || msg.Message == "alice" || msg.Message == "lis" {
		resp = getRandomNameCall(msg)
	} else if msg.Message == "Alif" || msg.Message == "lif" || msg.Message == "liff" || msg.Message == "lip" || msg.Message == "lipp" {
		resp = getRandomMasterCall(msg)
	} else if msg.Message == "p" {
		resp = getRandomRoast(msg)
	} else if ws.IsLoginRequest(msg, WAKeyword) {
		resp = HandlerQRLogin(msg, WAKeyword)
	} else {
		resp = getRandomGreetings(msg)
	}
	return c.JSON(resp)
}
