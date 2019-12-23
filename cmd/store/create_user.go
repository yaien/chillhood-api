package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/api/models"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
	"log"
)

func createUser() *cobra.Command {
	var name string
	var email string
	var password string

	cmd := &cobra.Command{
		Use: "users:create",
		Short: "Create an user that can admin the cloth store api",
		Run: func(cmd *cobra.Command, args []string){
			app, err := core.NewApp()
			if err != nil {
				log.Fatal(err)
			}
			service := services.NewUserService(app.DB)
			user := models.User{
				Role: "admin",
				Password: password,
				Email: email,
				Name: name,
			}
			user.HashPassword()
			service.Create(&user)
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
