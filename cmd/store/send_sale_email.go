package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"github.com/yaien/clothes-store-api/pkg/interface/mongodb"
	"github.com/yaien/clothes-store-api/pkg/service"
)

func sendSaleEmail() *cobra.Command {
	return &cobra.Command{
		Use:   "email:sale [invoice-ref]",
		Short: "send the sale email of a given invoice",
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
			emails.NotifySale(invoice)
			return nil
		},
	}
}
