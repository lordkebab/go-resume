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
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Resume struct {
	Header         ResumeHeader
	Experience     []ResumeExperience
	Education      []ResumeEducation
	Skills         []string
	Certifications []string
}

type ResumeHeader struct {
	Name    string
	Email   string
	Phone   string
	Website string
	AboutMe string
}

type ResumeExperience struct {
	CompanyName       string
	CompanyCity       string
	CompanyState      string
	CompanyStartMonth string
	CompanyStartYear  int
	CompanyEndMonth   string
	CompanyEndYear    int
	JobTitle          string
	JobDescription    string
	JobBullets        []string
	CurrentlyEmployed bool
}

type ResumeEducation struct {
	SchoolName         string
	SchoolCity         string
	SchoolState        string
	SchoolStartMonth   string
	SchoolStartYear    int
	SchoolEndMonth     string
	SchoolEndYear      int
	DegreeType         string
	Major              string
	MajorFocus         string
	Minor              string
	CurrentlyAttending bool
}

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a resume from a file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// get flag values
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		buf, err := ioutil.ReadFile(inputFile)
		checkError(err)

		// unmarshal the file into a data structure
		data := map[string]interface{}{}
		err = json.Unmarshal(buf, &data)
		if err != nil {
			// try yaml
			yaml.Unmarshal(buf, &data)
		}

		out, err := os.Create(outputFile)
		checkError(err)
		defer out.Close()

		t, err := template.ParseGlob("templates/*")
		checkError(err)

		// spit out the template
		err = t.Execute(out, data)
		checkError(err)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildCmd.Flags().StringP("input", "i", "resume.json", "Filename with your resume data")
	buildCmd.Flags().StringP("output", "o", "resume.html", "Name of the HTML file to output")

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
