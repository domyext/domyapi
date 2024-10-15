package webhook

import "os"

var WebhookSecret = os.Getenv("SECRET")

var WAAPIToken = os.Getenv("TOKEN")

const WAKeyword = "wh4t5auth0"
