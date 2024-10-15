package webhook

type Reply struct {
	Message string `bson:"messsage"`
}

type Header struct {
	Secret string `reqHeader:"secret"`
}
