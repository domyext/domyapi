package webhook

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/schollz/closestmatch"
	"github.com/whatsauth/wa"
)

var matcher = closestmatch.New([]string{"Babi", "Anjing", "goblok", "sayang", "syg", "cinta", "cantik", "Alice", "alice", "lis", "Alif", "lif", "liff", "lip", "lipp", "p"}, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1})

func HandlerIncomingMessage(msg model.IteungMessage) (resp atmessage.Response) {
	rand.Seed(time.Now().UnixNano())

	var messageTemplates []string

	// Use the closestmatch library to find the closest match
	closestMatch := matcher.Closest(msg.Message)

	switch closestMatch {
	case "Babi", "bgst" , "Anjing", "goblok", "tolol":
		messageTemplates = []string{
			"Ih sok asik",
			"Ampun bang jagooo",
			"Ih najissss",
			"Wah bocah bocah sok iye kaya gini nihh",
			"Ih lawak, apa lu?",
			"Ugh, get a life!",
			"Hahaha, you must be kidding.",
			"Jangan sok tahu deh!",
			"Boh, che cosa ridicola!",
			"Sei davvero così stupido?",
		}
	case "sayang", "syg", "cinta":
		messageTemplates = []string{
			"Haloo sayanggg, lagi ngapain?",
			"Eh.. aku tuh sayang kamu juga loooo",
			"Love you babee <3",
			"Apa sayangg kuu",
			"peluk dong yanggg",
		}
	case "cantik":
		messageTemplates = []string{
			"Makasih kaaaa, aku jadi maluuu ihh",
			"kaka jugaaa",
			"Love you kaaa <3",
		}
	case "Alice", "alice", "lis":
		messageTemplates = []string{
			"Oioioioi kenapa kaaa?",
			"Kaka manggil Alice? Ada apaa?",
			"maless.....",
			"Ada apa kak",
			"Ayoy kapten, napa?",
		}
	default:
		messageTemplates = []string{
			"Hai hai hai kak " + msg.Alias_name + "",
			"Hello " + msg.Alias_name + ", how are you?",
			"Hey there, " + msg.Alias_name + "! What's up?",
			// ... (other templates)
		}
	}

	// Randomly select a message template
	selectedTemplate := messageTemplates[rand.Intn(len(messageTemplates))]

	// Create the message with the selected template
	message := fmt.Sprintf(selectedTemplate)

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
