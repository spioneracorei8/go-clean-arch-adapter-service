package models

type MailForm struct {
	SenderName string `json:"sender_name"`
	To         string `json:"to"`
	ToName     string `json:"to_name"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
}
