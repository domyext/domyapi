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

	var messageTemplates []string

	if msg.Message == "Babi" || msg.Message == "Anjing" || msg.Message == "goblok" {
		messageTemplates = []string{
			"Ih sok asik",
			"Lah gw pikir itu emak lu!!",
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
	} else if msg.Message == "sayang" || msg.Message == "syg" || msg.Message == "cinta" {
		messageTemplates = []string{
			"Haloo sayanggg, lagi ngapain?",
			"Eh.. aku tuh sayang kamu juga loooo",
			"Love you babee <3",
			"Apa sayangg kuu",
			"peluk dong yanggg",
			"JOMBLO CARI PACAR SONO IHHH",
		}
	} else if msg.Message == "cantik" {
		messageTemplates = []string{
			"Makasih kaaaa, aku jadi maluuu ihh",
			"kaka jugaaa",
			"Love you kaaa <3",
		}
	} else if msg.Message == "Alice" || msg.Message == "alice" || msg.Message == "lis" {
		messageTemplates = []string{
			"Oioioioi kenapa kaaa?",
			"Kaka manggil Alice? Ada apaa?",
			"maless.....",
			"Ada apa kak",
			"Ayoy kapten, napa?",
		}
	} else if msg.Message == "Alif" || msg.Message == "lif" || msg.Message == "liff" || msg.Message == "lip" || msg.Message == "lipp" {
		messageTemplates = []string{
			"Ka Alif nya lagi turu, nanti aja yaaa",
			"Kata Alif, dia lagi sibuk. Nanti yaaa",
			"Alif nya lagi keluar, tunggu bentar yaaa",
			"Ada apa kak? Nanti aku kasioh tau Alif nya",
			"Halo kaka, Aku Alice botnya Alif. Ada yang bisa aku bantu?",
		}
	} else if msg.Message == "p" {
		messageTemplates = []string{
			"p p p p p p pant*",
			"Minimal salam dong bos.. pa pe pap pe palalo petak",
			"BENGEUT SIA PETAK ANJAY",
			"Salam yang bener bang",
			"Q",
			"Oh, trying to be clever again?",
			"You must have a Ph.D. in nonsense.",
			"Jangan bangga deh, coba pakai otak sekali-kali.",
			"Sei così divertente che sto per svenire...",
		}
	} else {
		messageTemplates = []string{
			"Hai hai hai kak " + msg.Alias_name + "",
			"Hello " + msg.Alias_name + ", how are you?",
			"Hey there, " + msg.Alias_name + "! What's up?",
			"Greetings, " + msg.Alias_name + "!",
			"Hi " + msg.Alias_name + ", nice to hear from you!",
			"Halo ka " + msg.Alias_name + ", lagi nyari apa nihh?",
			"Ciao " + msg.Alias_name + ", come va oggi?",
			"Bonjour " + msg.Alias_name + ", comment ça va?",
			"Hola " + msg.Alias_name + ", ¿cómo estás?",
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
