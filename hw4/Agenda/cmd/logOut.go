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

// logOutCmd represents the logOut command
var logOutCmd = &cobra.Command{
	Use:   "logOut",
	Short: "Log out",
	Long: `Use logOut to log out a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LogOut called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		var s myAgenda.Storage
		s.ReadFormFile()
		valid := s.QueryUser(func(user myAgenda.User) bool {
			return user.M_name == username && user.M_password == password
		})
		if valid.Len()==0{
			fmt.Fprintf(os.Stderr, "Log out failed!\n")
			return
		}

		curValid := s.QueryCurUser(func(s string) bool {
			return s == username
		})

		if curValid.Len()==0{
			fmt.Fprintf(os.Stderr, "LogOut fail!: User not logged in!\n")
			return
		}

		s.DeleteCurUser(username)
		s.WriteToFile()
		fmt.Println("LogOut successed")
	},
}

func init() {
	rootCmd.AddCommand(logOutCmd)
	logOutCmd.Flags().StringP("username", "u", "username", "Username")
	logOutCmd.Flags().StringP("password", "p", "password", "Password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logOutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logOutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
