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

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _:= cmd.Flags().GetString("id")
		currentTime := time.Now()
		if id == "" {
			panic("please provide id of task ");
		}
		db := dbConn()
		insForm, err := db.Exec("UPDATE tasks SET status=? , completedAt=? WHERE id=?", "1", currentTime, id)
        if err != nil {
            panic(err.Error())
			fmt.Println(insForm);
        }
        

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
			fmt.Println(`You have completed the "`+taskname+`" task.`)
			
		}
		
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
	doCmd.Flags().StringP("id", "i", "", "id of the task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
