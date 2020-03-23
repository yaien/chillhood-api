package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
	"github.com/yaien/clothes-store-api/pkg/core"
)

func createUser() *cobra.Command {
	var name string
	var email string
	var password string

	cmd := &cobra.Command{
		Use:   "users:create",
		Short: "Create an user that can admin the cloth store api",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := core.NewApp()
			if err != nil {
				log.Fatal(err)
			}
			service := services.NewUserService(app.DB)
			user := models.User{
				Role:     "admin",
				Password: password,
				Email:    email,
				Name:     name,
			}
			user.HashPassword()
			if err := service.Create(&user); err != nil {
				log.Fatal(err)
			}
			bytes, _ := json.MarshalIndent(&user, "", "    ")
			fmt.Println(string(bytes))
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&name, "name", "", "")
	flags.StringVar(&email, "email", "", "")
	flags.StringVar(&password, "password", "", "")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("password")

	return cmd
}
