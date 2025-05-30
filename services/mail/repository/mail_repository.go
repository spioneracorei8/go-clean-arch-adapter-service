package repository

import (
	"adapter-service/models"
	my_mail "adapter-service/services/mail"
	"context"

	"gopkg.in/mail.v2"
)

type mailRepository struct {
}

func NewMailRepositoryImpl() my_mail.MailRepository {
	return &mailRepository{}
}

func (m *mailRepository) SendMail(ctx context.Context, form *models.MailForm) error {
	email := mail.NewMessage()
	email.SetHeader("From", form.SenderName)
	email.SetHeader("To", form.To)
	email.SetHeader("Subject", form.Subject)
	email.SetBody("text/html", form.Body)

	d := mail.NewDialer("", 0, "", "")
	if err := d.DialAndSend(email); err != nil {
		return err
	}

	// msg, _ := json.Marshal(job)
	// if err := e.queueProducer.PublishMessage(string(msg)); err != nil {

	// }

	// msg, _ := json.Marshal(job)
	// if err := e.queueProducer.PublishMessage(string(msg)); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)

	return nil
}
