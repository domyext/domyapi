package webhook

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
)

func getRandomGreetings(msg model.IteungMessage) (resp atmessage.Response) {
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
		"Ciao %s, come va oggi?",
		"Bonjour %s, comment ça va?",
		"Hola %s, ¿cómo estás?",
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

	return resp

}

func getRandomRude(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
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

	return resp
}

func getRandomSadBoi(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
		"Haloo sayanggg, lagi ngapain?",
		"Eh.. aku tuh sayang kamu juga loooo",
		"Love you babee <3",
		"Apa sayangg kuu",
		"peluk dong yanggg",
		"JOMBLO CARI PACAR SONO IHHH",
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

	return resp
}

func getRandomCompliment(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
		"Makasih kaaaa, aku jadi maluuu ihh",
		"kaka jugaaa",
		"Love you kaaa <3",
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

	return resp
}

func getRandomNameCall(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
		"Oioioioi kenapa kaaa?",
		"Kaka manggil Alice? Ada apaa?",
		"maless.....",
		"Ada apa kak",
		"Ayoy kapten, napa?",
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

	return resp
}

func getRandomMasterCall(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
		"Ka Alif nya lagi turu, nanti aja yaaa",
		"Kata Alif, dia lagi sibuk. Nanti yaaa",
		"Alif nya lagi keluar, tunggu bentar yaaa",
		"Ada apa kak? Nanti aku kasioh tau Alif nya",
		"Halo kaka, Aku Alice botnya Alif. Ada yang bisa aku bantu?",
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

	return resp
}

func getRandomRoast(msg model.IteungMessage) (resp atmessage.Response) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of message templates
	messageTemplates := []string{
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

	return resp
}

// getRandomMessage returns a random message from the provided slice of messages.
func randomizeMsg(messages []string) string {
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(len(messages))]
}

// Randomize response if secret code was invalid
func getRandomIncorrectSecretMessage() string {
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

	return randomizeMsg(incorrectSecretMessages)
}
