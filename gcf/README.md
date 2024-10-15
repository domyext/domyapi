# Webhook Layanan WhatsAuth
WebHook Menggunakan method HTTP POST dengan Header bernama Secret  
![image](https://github.com/whatsauth/webhook/assets/11188109/7734295e-89bb-4b05-ab05-d2ee0bdb6019)  
Format JSON yang dikirim ke WebHook menggunakan struct [IteungMessage](https://github.com/aiteung/module/blob/11c330c30b86f7e7c04bf3d60f04b928c90fd6dc/model/main.go#L23) :
```go
type WAMessage struct {
	Phone_number       string  `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Reply_phone_number string  `json:"reply_phone_number,omitempty" bson:"reply_phone_number,omitempty"`
	Chat_number        string  `json:"chat_number,omitempty" bson:"chat_number,omitempty"`
	Chat_server        string  `json:"chat_server,omitempty" bson:"chat_server,omitempty"`
	Group_name         string  `json:"group_name,omitempty" bson:"group_name,omitempty"`
	Group_id           string  `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Group              string  `json:"group,omitempty" bson:"group,omitempty"`
	Alias_name         string  `json:"alias_name,omitempty" bson:"alias_name,omitempty"`
	Message            string  `json:"messages,omitempty" bson:"messages,omitempty"`
	From_link          bool    `json:"from_link,omitempty" bson:"from_link,omitempty"`
	From_link_delay    uint32  `json:"from_link_delay,omitempty" bson:"from_link_delay,omitempty"`
	Is_group           bool    `json:"is_group,omitempty" bson:"is_group,omitempty"`
	Filename           string  `json:"filename,omitempty" bson:"filename,omitempty"`
	Filedata           string  `json:"filedata,omitempty" bson:"filedata,omitempty"`
	Latitude           float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude          float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
	LiveLoc            bool    `json:"liveloc,omitempty" bson:"liveloc,omitempty"`
}
```
[Contoh Source Code WebHook](./function.go)


## Langkah-langkah pembuatan GCP
Pastikan sudah setting environment variabel pada GCF antara lain :
1. SECRET : sebagai pengaman endpoint GCF webhook yang di tembak dari whatsauth
2. TOKEN : token yang digunakan untuk menggunakan API WhatsAuth
3. Buat Cloud Function dan pilih server jakarta indonesia  
   ![image](https://github.com/whatsauth/webhook/assets/11188109/ad72a002-b318-4475-8c85-94b266aaa4a5)
4. Masukkan variabel environtment SECRET dan TOKEN. Harap diingat ini case sensitif pastikan huruf besar semua untuk Name dari variabel.  
   ![image](https://github.com/whatsauth/webhook/assets/11188109/5ce519e9-c9ee-45aa-ad58-edd14a4c661d)
5. Ambil kode program GCF [disini](https://github.com/whatsauth/webhook/blob/main/gcf/function.go), pastikan Entry Point sama dengan yang ada di init.  
   ![image](https://github.com/whatsauth/webhook/assets/11188109/fd30ddad-eca0-452a-8e1d-d7038401f7e6)
6. Agar API bisa diakses Publik Berikan Akses Publik dari Menu Cloud Run dan Tambahkan permission sebagai berikut:  
   New Principals : AllUsers  
   Role : Cloud Run Invoker  
   ![image](https://github.com/petapedia/gcf/assets/11188109/45b84091-3e42-4124-9264-ae86e08d49b3)  


## Release Package
```sh
go get -u all					#update existing package
go mod tidy					#generate go mod
git tag                                 	#check current version
git tag v0.0.3                          	#set tag version
git push origin --tags                  	#push tag version to repo
go list -m github.com/whatsauth/webhook@v0.0.3   #publish to pkg dev, replace ORG/URL with your repo URL
```
