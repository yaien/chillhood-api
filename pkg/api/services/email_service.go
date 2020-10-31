package services

import (
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jordan-wright/email"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/core"
	"html/template"
	"log"
	"net/smtp"
	"strings"
)

type EmailService interface {
	NotifySale(invoice *models.Invoice)
	NotifyTransport(invoice *models.Invoice)
}

type emailService struct {
	config    *core.SMTPConfig
	templates *core.Templates
}

func (e *emailService) functions() template.FuncMap {
	return template.FuncMap{
		"first": func(name string) string {
			return strings.Split(name, " ")[0]
		},
		"link": func(ref string) string {
			return strings.ReplaceAll(e.config.RefLink, "{ref}", ref)
		},
		"currency": func(value int) string {
			return fmt.Sprintf("$%s", humanize.Comma(int64(value)))
		},
	}
}

func (e *emailService) NotifySale(invoice *models.Invoice) {
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

func (e *emailService) NotifyTransport(invoice *models.Invoice) {
	var buffer bytes.Buffer
	err := e.templates.Transport.Execute(&buffer, invoice)
	if err != nil {
		log.Printf("failed execiting sale template for invoice #%s: %s", invoice.Ref, err.Error())
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
		log.Printf("failed sending sale email for invoice #%s: %s", invoice.Ref, err.Error())
	}
}

func NewEmailService(config *core.SMTPConfig, templates *core.Templates) EmailService {
	return &emailService{config, templates}
}
