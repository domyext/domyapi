package webhook

import (
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/ws"
)

func HandlerQRLogin(msg model.IteungMessage, WAKeyword string) (resp atmessage.Response) {
	dt := &ws.WhatsauthRequest{
		Uuid:        ws.GetUUID(msg, WAKeyword),
		Phonenumber: msg.Phone_number,
		Delay:       msg.From_link_delay,
	}
	resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", WAAPIToken, dt, "https://api.wa.my.id/api/whatsauth/request")
	return
}
