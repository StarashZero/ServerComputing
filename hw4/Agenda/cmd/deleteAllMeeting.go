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

// deleteAllMeetingCmd represents the deleteAllMeeting command
var deleteAllMeetingCmd = &cobra.Command{
	Use:   "deleteAllMeeting",
	Short: "Delete all meetings of a sponsor",
	Long: `Use deleteAllMeeting to delete all meetings of a sponsor
sponsor need to be logged in`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteAllMeeting called")
		sponsor, _ := cmd.Flags().GetString("sponsor")
		var s myAgenda.Storage
		s.ReadFormFile()

		valid := s.QueryCurUser(func(username string) bool {
			return username == sponsor
		})

		if valid.Len() == 0 {
			fmt.Fprintf(os.Stderr, "DeleteAllMeeting failed!: Sponsor not log in!\n")
			return
		}

		cnt:=s.DeleteMeeting(func(meeting myAgenda.Meeting) bool {
			return meeting.M_sponsor == sponsor
		})

		if cnt==0{
			fmt.Fprintf(os.Stderr, "DeleteAllMeeting failed!: No matching meeting for username!\n")
			return
		}

		s.WriteToFile()
		fmt.Println("DeleteAllMeeting successed")
	},
}

func init() {
	rootCmd.AddCommand(deleteAllMeetingCmd)
	deleteAllMeetingCmd.Flags().StringP("sponsor", "u", "sponsor", "Sponsor")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteAllMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteAllMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
