package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
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

func main() {

	buf, err := ioutil.ReadFile("resume.json")
	checkError(err)

	// unmarshal the file into a data structure
	data := map[string]interface{}{}
	json.Unmarshal(buf, &buf)

	out, err := os.Create("index.html")
	checkError(err)
	defer out.Close()

	t, err := template.ParseGlob("templates/*")
	checkError(err)

	// spit out the template
	err = t.Execute(out, data)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
