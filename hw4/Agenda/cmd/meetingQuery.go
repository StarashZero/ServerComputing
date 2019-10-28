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

// meetingQueryCmd represents the meetingQuery command
var meetingQueryCmd = &cobra.Command{
	Use:   "meetingQuery",
	Short: "Query a meeting",
	Long: `Use meetingQuery to query a meeting
username need to be logged in`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MeetingQuery called")
		username, _ := cmd.Flags().GetString("username")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")

		var s myAgenda.Storage
		s.ReadFormFile()

		valid := s.QueryCurUser(func(user string) bool {
			return user == username
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "MeetingQuery failed!: User not log in!\n")
			return
		}

		d1, d2 := myAgenda.StringToDate(startTime), myAgenda.StringToDate(endTime)

		if myAgenda.IsValid(d1)&&myAgenda.IsValid(d2){
			valid := s.QueryMeeting(func(meeting myAgenda.Meeting) bool {
				return (meeting.M_sponsor==username||meeting.IsParticipator(username))&&!(myAgenda.CompareDate(d2,meeting.M_startDate)<0||myAgenda.CompareDate(d1,meeting.M_endDate)>0)
			})

			for i:=valid.Front();i!=nil;i=i.Next(){
				met := i.Value.(myAgenda.Meeting)
				fmt.Printf("Title: %s\tSponser: %s\tStart Time: %s\tEnd Time: %s\tParticipators: %s\n", met.M_title,met.M_sponsor, met.M_startDate.DateToString(), met.M_endDate.DateToString(), met.M_participators)
			}
		}

		fmt.Println("MeetingQuery successed")
	},
}

func init() {
	rootCmd.AddCommand(meetingQueryCmd)
	meetingQueryCmd.Flags().StringP("username", "u", "username", "Username")
	meetingQueryCmd.Flags().StringP("startTime", "s", "start time", "Start time(yyyy-mm-dd/hh:mm)")
	meetingQueryCmd.Flags().StringP("endTime", "e", "end time", "End time(yyyy-mm-dd/hh:mm)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meetingQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meetingQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
