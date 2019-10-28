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

// addMeetingParticipatorCmd represents the addMeetingParticipator command
var addMeetingParticipatorCmd = &cobra.Command{
	Use:   "addMeetingParticipator",
	Short: "Add a participator to a meeting",
	Long: `Use addMeetingParticipator to add a participator to a meeting
The sponsor need to be logged in and participator need to be registered`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AddMeetingParticipator called")
		sponsor, _ := cmd.Flags().GetString("sponsor")
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("participator")

		var s myAgenda.Storage
		s.ReadFormFile()

		if sponsor == participator {
			fmt.Fprintf(os.Stderr, "AddMeetingParticipator failed!: The sponsor can't be the participator!\n")
			return
		}

		valid := s.QueryCurUser(func(username string) bool {
			return username == sponsor
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "AddMeetingParticipator failed!: Sponsor not log in!\n")
			return
		}

		valid = s.QueryUser(func(user myAgenda.User) bool {
			return user.M_name == participator
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "AddMeetingParticipator failed!: Participator isn't in user!\n")
			return
		}

		cnt := s.UpdateMeeting(func(meeting myAgenda.Meeting) bool {
			if meeting.M_title == title && meeting.M_sponsor == sponsor && !meeting.IsParticipator(participator) {
				valid = s.QueryMeeting(func(meeting2 myAgenda.Meeting) bool {
					return (!(myAgenda.CompareDate(meeting2.M_startDate, meeting.M_endDate) >= 0 || myAgenda.CompareDate(meeting2.M_endDate, meeting.M_startDate) <= 0)) && (meeting2.IsParticipator(participator) || meeting2.M_sponsor == participator)
				})
				return valid.Len() == 0
			}
			return false
		}, func(meeting *myAgenda.Meeting) {
			meeting.AddParticipator(participator)
		})

		if cnt == 0 {
			fmt.Fprintf(os.Stderr, "AddMeetingParticipator failed!: Conflicting meeting of participator!(date or title)\n")
			return
		}

		s.WriteToFile()
		fmt.Println("AddMeetingParticipator successed")
	},
}

func init() {
	rootCmd.AddCommand(addMeetingParticipatorCmd)
	addMeetingParticipatorCmd.Flags().StringP("sponsor", "u", "sponsor", "Sponsor of meeting")
	addMeetingParticipatorCmd.Flags().StringP("title", "t", "title", "Title of meeting")
	addMeetingParticipatorCmd.Flags().StringP("participator", "p", "participator", "participator need to be added")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addMeetingParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addMeetingParticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
