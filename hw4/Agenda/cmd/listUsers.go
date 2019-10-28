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

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "listUsers",
	Short: "List all users registered",
	Long: `Use listUsers to list all users registered
username needed to logged in`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListUsers called")
		username, _ := cmd.Flags().GetString("username")

		var s myAgenda.Storage
		s.ReadFormFile()

		curValid := s.QueryCurUser(func(s string) bool {
			return s == username
		})

		if curValid.Len()==0{
			fmt.Fprintf(os.Stderr, "ListUsers failed!: User not log in!\n")
			return
		}

		valid := s.QueryUser(func(user myAgenda.User) bool {
			return true
		})

		for i:=valid.Front();i!=nil;i=i.Next(){
			fmt.Printf("Username: %s\tEmail: %s\tPhone: %s\n", i.Value.(myAgenda.User).M_name, i.Value.(myAgenda.User).M_email, i.Value.(myAgenda.User).M_phone)
		}

		fmt.Println("ListUsers successed")
	},
}

func init() {
	rootCmd.AddCommand(listUsersCmd)
	listUsersCmd.Flags().StringP("username", "u", "username", "username")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listUsersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listUsersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
