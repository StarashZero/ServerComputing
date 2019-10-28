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

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "Quit a meeting",
	Long: `Use quitMeeting to quit a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("QuitMeeting called")
		username, _ := cmd.Flags().GetString("username")
		title, _ := cmd.Flags().GetString("title")
		var s myAgenda.Storage
		s.ReadFormFile()

		valid := s.QueryCurUser(func(user string) bool {
			return user == username
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "QuitMeeting failed!: User not log in!\n")
			return
		}

		cnt:=s.UpdateMeeting(func(meeting myAgenda.Meeting) bool {
			return meeting.M_title==title&&meeting.IsParticipator(username)&&meeting.M_sponsor!=username
		}, func(meeting *myAgenda.Meeting) {
			meeting.RemoveParticipator(username)
		})

		if cnt==0{
			fmt.Fprintf(os.Stderr, "QuitMeeting failed!: No matching meeting for username or username is the sponsor!\n")
			return
		}

		s.DeleteMeeting(func(meeting myAgenda.Meeting) bool {
			return len(meeting.M_participators) == 0
		})

		s.WriteToFile()
		fmt.Println("QuitMeeting successed")
	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)
	quitMeetingCmd.Flags().StringP("username", "u", "username", "Username")
	quitMeetingCmd.Flags().StringP("title", "t", "title", "Title of meeting")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
