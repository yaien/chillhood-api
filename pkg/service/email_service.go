package service

import (
	"bytes"
	"github.com/jordan-wright/email"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"log"
	"net/smtp"
)

type EmailService interface {
	NotifySale(invoice *entity.Invoice)
	NotifyTransport(invoice *entity.Invoice)
}

type emailService struct {
	config    *infrastructure.SMTPConfig
	templates *infrastructure.Templates
}

func (e *emailService) NotifySale(invoice *entity.Invoice) {
	var buffer bytes.Buffer
	err := e.templates.Sale.Execute(&buffer, invoice)
	if err != nil {
		log.Printf("failed execiting sale template for invoice #%s: %s", invoice.Ref, err.Error())
		return
	}
	m := &email.Email{
		ReplyTo: nil,
		From:    e.config.Sender,
		To:      []string{invoice.Shipping.Email},
		Subject: "¡Gracias por comprar en Chillhood!",
		HTML:    buffer.Bytes(),
	}

	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.Host)
	err = m.Send(e.config.Address(), auth)
	if err != nil {
		log.Printf("failed sending sale email for invoice #%s: %s", invoice.Ref, err.Error())
	}
}

func (e *emailService) NotifyTransport(invoice *entity.Invoice) {
	var buffer bytes.Buffer
	err := e.templates.Transport.Execute(&buffer, invoice)
	if err != nil {
		log.Printf("failed executing sale template for invoice #%s: %s", invoice.Ref, err.Error())
		return
	}
	m := &email.Email{
		ReplyTo: nil,
		From:    e.config.Sender,
		To:      []string{invoice.Shipping.Email},
		Subject: "¡Tu pedido esta en camino!",
		HTML:    buffer.Bytes(),
	}

	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.Host)
	err = m.Send(e.config.Address(), auth)
	if err != nil {
		log.Printf("failed sending sale email for invoice #%s: %s", invoice.Ref, err.Error())
	}
}

func NewEmailService(config *infrastructure.SMTPConfig, templates *infrastructure.Templates) EmailService {
	return &emailService{config, templates}
}
