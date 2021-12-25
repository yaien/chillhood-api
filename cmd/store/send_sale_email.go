package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/api/services"
	"github.com/yaien/clothes-store-api/pkg/core"
	"github.com/yaien/clothes-store-api/pkg/interface/repository"
)

func sendSaleEmail() *cobra.Command {
	return &cobra.Command{
		Use:   "email:sale [invoice-ref]",
		Short: "send the sale email of a given invoice",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := core.NewApp()
			if err != nil {
				return err
			}
			invoices := services.NewInvoiceService(repository.NewMongoInvoiceRepository(app.DB))
			emails := services.NewEmailService(app.Config.SMTP, app.Templates)
			invoice, err := invoices.FindOneByRef(context.TODO(), args[0])
			if err != nil {
				return fmt.Errorf("failed finding invoice: %w", err)
			}
			emails.NotifySale(invoice)
			return nil
		},
	}
}
