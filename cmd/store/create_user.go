package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/interface/mongodb"
	"github.com/yaien/clothes-store-api/pkg/service"
	"log"

	"github.com/spf13/cobra"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
)

func createUser() *cobra.Command {
	var name string
	var email string
	var password string

	cmd := &cobra.Command{
		Use:   "users:create",
		Short: "Create an user that can admin the cloth store rest",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := infrastructure.NewApp()
			if err != nil {
				log.Fatal(err)
			}
			srv := service.NewUserService(mongodb.NewUserRepository(app.DB))
			user := entity.User{
				Role:     "admin",
				Password: password,
				Email:    email,
				Name:     name,
			}
			user.HashPassword()
			if err := srv.Create(context.TODO(), &user); err != nil {
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
