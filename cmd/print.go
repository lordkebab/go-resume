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
	"io/ioutil"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the resume schema",
	Long:  `Output a blank resume schema to stdout or a file`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string

		// figure out the filename of the schema we will read in
		format, _ := cmd.Flags().GetString("format")
		if format == "json" {
			fileName = "example.json"
		} else if format == "yaml" || format == "yml" {
			fileName = "example.yml"
		} else {
			// raise error and quit
			fmt.Println("error")
		}

		// read in the schema
		schema, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", fileName, err)
		}

		output, _ := cmd.Flags().GetString("output")
		stdout, _ := cmd.Flags().GetBool("stdout")

		if stdout == true {
			fmt.Printf("%s\n", schema)
		} else {
			if output == "" {
				fmt.Println("Value for output is required if not printing to stdout")
			} else {
				ioutil.WriteFile(output, schema, 0644)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	printCmd.Flags().BoolP("stdout", "s", false, "Print the resume schema to stdout")
	printCmd.Flags().StringP("output", "o", "", "Filename to save the resume schema")
	printCmd.Flags().StringP("format", "f", "json", "Format for the resume schema (json or yaml)")
}
