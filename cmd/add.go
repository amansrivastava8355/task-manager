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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `Add a new task to your TODO list`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("string")

		if name == "" {
			panic("please provide task title")
		}
		db := dbConn()

		stmt, err := db.Prepare("INSERT INTO tasks(taskname, status) VALUES(?, ?);")
		if err != nil {
			panic(err.Error())
		}
		stmt.Exec(name, 0)

		fmt.Println(`Added "` + name + `" to your task list.`)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("string", "s", "", "Add a new task to your TODO list")
}
