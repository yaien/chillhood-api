package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"github.com/yaien/clothes-store-api/pkg/interface/mongodb"
	"github.com/yaien/clothes-store-api/pkg/service"
)

func sendTransportEmail() *cobra.Command {
	return &cobra.Command{
		Use:   "email:transport [invoice-ref]",
		Short: "send the transport email of a given invoice",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := infrastructure.NewApp()
			if err != nil {
				return err
			}
			invoices := service.NewInvoiceService(mongodb.NewMongoInvoiceRepository(app.DB))
			emails := service.NewEmailService(app.Config.SMTP, app.Templates)
			invoice, err := invoices.FindOneByRef(context.TODO(), args[0])
			if err != nil {
				return fmt.Errorf("failed finding invoice: %w", err)
			}
			if invoice.Status != entity.Completed {
				return errors.New("invoice status is not completed")
			}
			emails.NotifyTransport(invoice)
			return nil
		},
	}
}
