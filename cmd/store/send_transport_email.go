package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
	"github.com/yaien/clothes-store-api/pkg/core"
)

func sendTransportEmail() *cobra.Command {
	return &cobra.Command{
		Use:   "email:transport [invoice-ref]",
		Short: "send the transport email of a given invoice",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := core.NewApp()
			if err != nil {
				return err
			}
			invoices := services.NewInvoiceService(app.DB)
			emails := services.NewEmailService(app.Config.SMTP, app.Templates)
			invoice, err := invoices.GetByRef(args[0])
			if err != nil {
				return fmt.Errorf("failed finding invoice: %w", err)
			}
			if invoice.Status != models.Completed {
				return errors.New("invoice status is not completed")
			}
			emails.NotifyTransport(invoice)
			return nil
		},
	}
}
