package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/whatsauth/webhook"
)

func main() {
	http.HandleFunc("/", adaptor.FiberHandlerFunc(webhook.PostMessage))
	http.ListenAndServe(":8080", nil)
}
