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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "Create a meeting",
	Long: `Use createMeeting to create a meeting
sponsor need to be logged in
participatorNumber has to more than zero`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateMeeting called")
		sponsor, _ := cmd.Flags().GetString("sponsor")
		title, _ := cmd.Flags().GetString("title")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		participatorNumberTime, _ := cmd.Flags().GetInt("participatorNumber")

		var s myAgenda.Storage
		s.ReadFormFile()
		valid := s.QueryCurUser(func(username string) bool {
			return username == sponsor
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "Create Meeting failed!: Sponsor not log in!\n")
			return
		}

		d1, d2 := myAgenda.StringToDate(startTime), myAgenda.StringToDate(endTime)

		if !(myAgenda.IsValid(d1) && myAgenda.IsValid(d2)) {
			fmt.Fprintf(os.Stderr, "Create Meeting failed!: StartDate or endDate isn't valid!\n")
			return
		}

		if myAgenda.CompareDate(d1, d2) != -1 {
			fmt.Fprintf(os.Stderr, "Create Meeting failed!: StartDate more than or equal to endDate!\n")
			return
		}

		if participatorNumberTime <= 0 {
			fmt.Fprintf(os.Stderr, "Create Meeting failed!: The num of participator less than or equal zero!\n")
			return
		}

		valid = s.QueryMeeting(func(meeting myAgenda.Meeting) bool {
			return meeting.M_title == title
		})

		if valid.Len() != 0 {
			fmt.Fprintf(os.Stderr, "Create Meeting failed!: Title has existed!\n")
			return
		}

		var participators []string
		for i := 0; i < participatorNumberTime; i++ {
			fmt.Printf("Please enter NO.%d participator: ", i+1)
			var participator string
			fmt.Scanf("%s\n", &participator)
			participators = append(participators, participator)
		}

		valid = s.QueryMeeting(func(meeting myAgenda.Meeting) bool {
			return (!((myAgenda.CompareDate(meeting.M_startDate, d2) >= 0) || (myAgenda.CompareDate(meeting.M_endDate, d1) <= 0))) && (meeting.IsParticipator(sponsor) || (meeting.M_sponsor == sponsor))
		})

		if valid.Len() != 0 {
			fmt.Fprintf(os.Stderr, "Create Meeting failed!: Conflicting meeting of sponsor!(date)\n")
			return
		}

		for i := 0; i < len(participators); i++ {
			if participators[i] == sponsor {
				fmt.Fprintf(os.Stderr, "Create Meeting failed!: The sponsor can't be the participator!\n")
				return
			}

			valid := s.QueryUser(func(user myAgenda.User) bool {
				return user.M_name == participators[i]
			})

			if valid.Len() == 0 {
				fmt.Fprintf(os.Stderr, "Create Meeting failed!: Participator isn't in user!\n")
				return
			}

			valid = s.QueryMeeting(func(meeting myAgenda.Meeting) bool {
				return (!((myAgenda.CompareDate(meeting.M_startDate, d2) >= 0) || (myAgenda.CompareDate(meeting.M_endDate, d1) <= 0))) && (meeting.IsParticipator(participators[i]) || (meeting.M_sponsor == participators[i]))
			})

			if valid.Len() != 0 {
				fmt.Fprintf(os.Stderr, "Create Meeting failed!: Conflicting meeting of participator!(date)\n")
				return
			}

			for j := i+1; j < len(participators); j++ {
				if participators[i] == participators[j] {
					fmt.Fprintf(os.Stderr, "Create Meeting failed!: Multiple participator!\n")
					return
				}
			}
		}

		meeting := myAgenda.Meeting{
			sponsor,
			participators,
			d1,
			d2,
			title,
		}

		s.CreateMeeting(meeting)
		s.WriteToFile()
		fmt.Println("CreateMeeting successed")
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringP("sponsor", "u", "sponsor", "Sponsor of meeting")
	createMeetingCmd.Flags().StringP("title", "t", "title", "Title of meeting")
	createMeetingCmd.Flags().StringP("startTime", "s", "start time", "Start time of meeting(yyyy-mm-dd/hh:mm)")
	createMeetingCmd.Flags().StringP("endTime", "e", "end time", "End time of meeting(yyyy-mm-dd/hh:mm)")
	createMeetingCmd.Flags().IntP("participatorNumber", "p", 0, "Number of participator")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
