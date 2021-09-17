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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "delete your task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")

		if id == "" {
			panic("please provide id of task ")
		}
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM tasks WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}
		for selDB.Next() {
			var task Task
			err = selDB.Scan(&task.id, &task.taskname, &task.status, &task.completedAt)
			if err != nil {
				panic(err.Error())
			}

			taskname := task.taskname
			fmt.Println(`You have deleted the "` + taskname + `" task.`)

		}
		delForm, err := db.Prepare("DELETE FROM tasks WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().StringP("id", "i", "", "id of the task")
}
