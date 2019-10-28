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

// removeMeetingParticipatorCmd represents the removeMeetingParticipator command
var removeMeetingParticipatorCmd = &cobra.Command{
	Use:   "removeMeetingParticipator",
	Short: "Remove a participator from a meeting",
	Long: `Use removeMeetingParticipator to remove a participator from a meeting
sponsor need to be logged in`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RemoveMeetingParticipator called")

		sponsor, _ := cmd.Flags().GetString("sponsor")
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("participator")

		var s myAgenda.Storage
		s.ReadFormFile()

		valid := s.QueryCurUser(func(username string) bool {
			return username == sponsor
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "RemoveMeetingParticipator failed!: Sponsor not log in!\n")
			return
		}

		cnt:=s.UpdateMeeting(func(meeting myAgenda.Meeting) bool {
			return (meeting.M_title==title)&&(meeting.M_sponsor==sponsor)&&(meeting.IsParticipator(participator))
		}, func(meeting *myAgenda.Meeting) {
			meeting.RemoveParticipator(participator)
		})

		if cnt==0{
			fmt.Fprintf(os.Stderr, "RemoveMeetingParticipator failed!: No matching meeting or participator!\n")
			return
		}

		s.DeleteMeeting(func(meeting myAgenda.Meeting) bool {
			return len(meeting.M_participators) == 0
		})

		s.WriteToFile()

		fmt.Println("RemoveMeetingParticipator successed")
	},
}

func init() {
	rootCmd.AddCommand(removeMeetingParticipatorCmd)
	removeMeetingParticipatorCmd.Flags().StringP("sponsor", "u", "sponsor", "Sponsor of meeting")
	removeMeetingParticipatorCmd.Flags().StringP("title", "t", "title", "Title of meeting")
	removeMeetingParticipatorCmd.Flags().StringP("participator", "p", "participator", "Participator need to be removed")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeMeetingParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeMeetingParticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
