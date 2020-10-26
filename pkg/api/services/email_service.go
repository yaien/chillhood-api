package services

import (
	"bytes"
	"github.com/jordan-wright/email"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/core"
	"log"
	"net/smtp"
)

type EmailService interface {
	NotifySale(invoice *models.Invoice)
	NotifyTransport(invoice *models.Invoice)
}

type emailService struct {
	config    *core.SMTPConfig
	templates *core.Templates
}

func (e *emailService) NotifySale(invoice *models.Invoice) {
	var buffer bytes.Buffer
	err := e.templates.Sale.Execute(&buffer, invoice)
	if err != nil {
		log.Printf("failed execiting sale template for invoice #%s: %w", invoice.Ref, err)
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
		log.Printf("failed sending sale email for invoice #%s: %w", invoice.Ref, err)
	}
}

func (e *emailService) NotifyTransport(invoice *models.Invoice) {
	var buffer bytes.Buffer
	err := e.templates.Transport.Execute(&buffer, invoice)
	if err != nil {
		log.Printf("failed execiting sale template for invoice #%s: %w", invoice.Ref, err)
		return
	}
	m := &email.Email{
		ReplyTo: nil,
		From:    e.config.Sender,
		To:      []string{invoice.Shipping.Email},
		Subject: "¡Tu compra esta en camino!",
		HTML:    buffer.Bytes(),
	}

	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.Host)
	err = m.Send(e.config.Address(), auth)
	if err != nil {
		log.Printf("failed sending sale email for invoice #%s: %w", invoice.Ref, err)
	}
}

func NewEmailService(config *core.SMTPConfig, templates *core.Templates) EmailService {
	return &emailService{config, templates}
}
