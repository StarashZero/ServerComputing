/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"../myAgenda"
	"github.com/spf13/cobra"
	"os"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Long: `Use register to register a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Register called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		var s myAgenda.Storage
		s.ReadFormFile()
		valid := s.QueryUser(func(user myAgenda.User) bool {
			return user.M_name == username
		})
		if valid.Len()!=0{
			fmt.Fprintf(os.Stderr, "Register fail!: User exited!\n")
			return
		}
		s.CreateUser(myAgenda.User{
			username,
			password,
			email,
			phone,
		})
		s.WriteToFile()
		fmt.Println("Register successed")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "username", "Username")
	registerCmd.Flags().StringP("password", "p", "password", "Password")
	registerCmd.Flags().StringP("email", "e", "email@123.com", "Email")
	registerCmd.Flags().StringP("phone", "t", "123456789", "Phone")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
