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
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "task",
		Short: "task is a CLI for managing your TODOs.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), args[0])
			return nil
		},
	}
}

func Test_main(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"hi-via-args"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi-via-args" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi-via-args", string(out))
	}
}

var in string

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Add",
		Short: "Add a new task to your TODO list",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), in)
			return nil
		},
	}
	cmd.Flags().StringVar(&in, "in", "", "This is a very important input.")
	return cmd
}

func emptyRun(*cobra.Command, []string) {}

// ArbitraryArgs never returns an error.
func ArbitraryArgs(cmd *cobra.Command, args []string) error {
	return nil
}
func Test_addCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "task", Run: emptyRun}
	childCmd := &cobra.Command{Use: "add", Args: ArbitraryArgs, Run: emptyRun}
	rootCmd.AddCommand(childCmd)

	_, err := executeCommand(rootCmd, "add", "legal", "args")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func Test_doCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "task", Run: emptyRun}
	childCmd := &cobra.Command{Use: "do", Args: ArbitraryArgs, Run: emptyRun}
	rootCmd.AddCommand(childCmd)

	_, err := executeCommand(rootCmd, "do", "legal", "args")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func Test_completedCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "task", Run: emptyRun}
	childCmd := &cobra.Command{Use: "completed", Args: ArbitraryArgs, Run: emptyRun}
	rootCmd.AddCommand(childCmd)

	_, err := executeCommand(rootCmd, "completed", "legal", "args")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func Test_listCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "task", Run: emptyRun}
	childCmd := &cobra.Command{Use: "list", Args: ArbitraryArgs, Run: emptyRun}
	rootCmd.AddCommand(childCmd)

	_, err := executeCommand(rootCmd, "list", "legal", "args")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
func Test_rmCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "task", Run: emptyRun}
	childCmd := &cobra.Command{Use: "rm", Args: ArbitraryArgs, Run: emptyRun}
	rootCmd.AddCommand(childCmd)

	_, err := executeCommand(rootCmd, "rm", "legal", "args")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
