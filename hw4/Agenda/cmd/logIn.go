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

// logInCmd represents the logIn command
var logInCmd = &cobra.Command{
	Use:   "logIn",
	Short: "Log in",
	Long: `Use logIn to log in to a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LogIn called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		var s myAgenda.Storage
		s.ReadFormFile()
		valid := s.QueryUser(func(user myAgenda.User) bool {
			return user.M_name == username && user.M_password == password
		})
		if valid.Len()==0{
			fmt.Fprintf(os.Stderr, "Log in failed!: username or password error\n")
			return
		}

		curValid := s.QueryCurUser(func(s string) bool {
			return s == username
		})

		if curValid.Len()!=0{
			fmt.Fprintf(os.Stderr, "Log in failed!: User has logged in!\n")
			return
		}

		s.CreateCurUser(username)
		s.WriteToFile()
		fmt.Println("Log in successed")
	},
}

func init() {
	rootCmd.AddCommand(logInCmd)
	logInCmd.Flags().StringP("username", "u", "username", "Username")
	logInCmd.Flags().StringP("password", "p", "password", "Password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logInCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logInCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
