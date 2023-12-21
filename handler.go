Copy code
package webhook

import (
	"math/rand"
	"time"

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

		// Handle login request
		if ws.IsLoginRequest(msg, WAKeyword) {
			resp = HandlerQRLogin(msg, WAKeyword)
		} else {
			resp = HandlerIncomingMessage(msg)
		}
	} else {
		// Random response for incorrect secret code
		resp.Response = getRandomIncorrectSecretMessage()
	}
	return c.JSON(resp)
}

// getRandomIncorrectSecretMessage returns a random message for incorrect secret code
func getRandomIncorrectSecretMessage() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of possible incorrect secret messages
	incorrectSecretMessages := []string{
		"Oops! It seems like the secret code is incorrect.",
		"Sorry, but the provided secret code is not valid.",
		"Hmm, something went wrong. Please check the secret code and try again.",
		"Error: Incorrect secret code. Please verify and resend your message.",
		"Ups! kode loginnya salah nihh.",
		"Maaf kak, kode kodenya ga valid :(",
		"Hmm, coba cek kode loginnya lagi deh...",
		"BANGG UDAH BANGG, KODENYA SALAHHH",
	}

	// Randomly select a message
	return incorrectSecretMessages[rand.Intn(len(incorrectSecretMessages))]
}