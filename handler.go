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
	if h.Secret == WebhookSecret {
		var msg model.IteungMessage
		err = c.BodyParser(&msg)
		if err != nil {
			return err
		}
		if ws.IsLoginRequest(msg, WAKeyword) { //untuk whatsauth request login
			resp = HandlerQRLogin(msg, WAKeyword)
		} else { //untuk membalas pesan masuk
			resp = HandlerIncomingMessage(msg)
		}
	} else {
		resp.Response = "Secret Salah"
	}
	return c.JSON(resp)
}
