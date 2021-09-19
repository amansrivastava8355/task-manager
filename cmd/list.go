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

	"github.com/spf13/cobra"
)

func listCommand(cmd *cobra.Command, args []string) {
	db := dbConn()

	selDB, err := db.Query("SELECT * FROM tasks where status=0")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("You have the following tasks:")
	for selDB.Next() {
		var task Task
		err = selDB.Scan(&task.id, &task.taskname, &task.status, &task.completedAt)
		if err != nil {
			panic(err.Error())
		}
		taskid := task.id
		taskname := task.taskname
		fmt.Println(taskid + ". " + taskname)
	}

}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listCommand(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
