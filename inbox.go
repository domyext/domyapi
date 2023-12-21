package webhook

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
)

func HandlerIncomingMessage(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
		"Hai hai hai kak %s",
		"Hello %s, how are you?",
		"Hey there, %s! What's up?",
		"Greetings, %s!",
		"Hi %s, nice to hear from you!",
		"Halo ka %s, lagi nyari apa nihh?",
		// Add more templates as needed
	}

	// Randomly select a message template
	selectedTemplate := messageTemplates[rand.Intn(len(messageTemplates))]

	// Create the message with the selected template
	message := fmt.Sprintf(selectedTemplate, msg.Alias_name)

	// Create a wa.TextMessage
	dt := &wa.TextMessage{
		To:       msg.Chat_number,
		IsGroup:  false,
		Messages: message,
	}

	// Check if the message is from a group
	if msg.Chat_server == "g.us" {
		dt.IsGroup = true
	}

	// Ignore messages from specific phone numbers
	if (msg.Phone_number != "628112000279") && (msg.Phone_number != "6283131895000") {
		// Post the message using atapi.PostStructWithToken
		resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", WAAPIToken, dt, "https://api.wa.my.id/api/send/message/text")
	}

	return resp
}
