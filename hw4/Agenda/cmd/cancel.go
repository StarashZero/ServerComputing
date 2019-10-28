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

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a user",
	Long: `Use cancel to cancel a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cancel called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		var s myAgenda.Storage
		s.ReadFormFile()
		valid := s.QueryUser(func(user myAgenda.User) bool {
			return user.M_name == username && user.M_password == password
		})
		if valid.Len()==0{
			fmt.Fprintf(os.Stderr, "Cancel failed!: Incorrect username or password!\n")
			return
		}

		s.DeleteMeeting(func(meeting myAgenda.Meeting) bool {
			return meeting.M_sponsor == username
		})

		s.UpdateMeeting(func(meeting myAgenda.Meeting) bool {
			return meeting.IsParticipator(username)
		}, func(meeting *myAgenda.Meeting) {
			meeting.RemoveParticipator(username)
		})

		s.DeleteMeeting(func(meeting myAgenda.Meeting) bool {
			return len(meeting.M_participators) == 0
		})

		curValid := s.QueryCurUser(func(s string) bool {
			return s == username
		})

		if curValid.Len()!=0{
			s.DeleteCurUser(username)
		}

		s.DeleteUser(valid.Front().Value.(myAgenda.User))
		s.WriteToFile()
		fmt.Println("Cancel successed")
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)
	cancelCmd.Flags().StringP("username", "u", "username", "username")
	cancelCmd.Flags().StringP("password", "p", "password", "password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
