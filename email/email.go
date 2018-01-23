package email

import (
	"encoding/json"

	"github.com/brunoleonel/email-consumer/model"
	"gopkg.in/gomail.v2"
)

var mailtrapUser = "<seu_usuario_do_mailtrap>"
var mailtrapPassword = "<sua_senha_do_mailtrap>"

//SendEmail - Função para envio de e-mail
func SendEmail(email model.Email) (err error) {

	message := gomail.NewMessage()
	message.SetHeader("From", "exemplo@mail.com")
	message.SetHeader("To", email.To)
	message.SetHeader("Subject", email.Subject)
	message.SetBody("text/plain", email.Message)

	dialer := gomail.NewDialer("smtp.mailtrap.io", 2525, mailtrapUser, mailtrapPassword)

	err = dialer.DialAndSend(message)

	return
}

//ParseMessage - Faz a conversão do json da fila para um Email
func ParseMessage(message []byte) (email model.Email, err error) {

	email = model.Email{}
	err = json.Unmarshal(message, &email)
	return
}
