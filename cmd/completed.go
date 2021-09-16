/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"time"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List of completed tasks",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// intervelhour, _:= cmd.Flags().GetString("intervel")
		// if intervelhour == "" {
		// 	intervelhour = "24";
		// }
		db := dbConn()
		currentTime := time.Now()
		selDB, err := db.Query("SELECT * FROM tasks where status=1 && created > ? - interval 24 hour;", currentTime)
		
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("You have finished the following tasks today:")
		for selDB.Next() {
			var task Task
			err = selDB.Scan(&task.id, &task.taskname, &task.status, &task.completedAt)
			if err != nil {
				panic(err.Error())
			}
			
			taskname := task.taskname
			fmt.Println("- "+taskname)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
