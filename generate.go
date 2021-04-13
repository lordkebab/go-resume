package main

import (
	"fmt"
	"html/template"
	"os"
)

type Resume struct {
	Header     ResumeHeader
	Experience []ResumeExperience
	Education  []ResumeEducation
	Skills     []ResumeSkills
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

type ResumeSkills struct {
	SkillName string
}

func main() {
	// tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := Resume{
		Header: ResumeHeader{
			Name:    "Matt Selph",
			Email:   "matt.selph@gmail.com",
			Phone:   "(478) 290-2599",
			Website: "https://matts.io",
			AboutMe: "Sr. Cloud Solutions Architect for the U.S. Air Force, focused on designing and implementing secure and highly-available cloud architecture.",
		},
		Experience: []ResumeExperience{
			{
				CompanyName:       "Booz Allen Hamilton",
				CompanyCity:       "Warner Robins",
				CompanyState:      "GA",
				CompanyStartMonth: "July",
				CompanyStartYear:  2019,
				CurrentlyEmployed: true,
				JobTitle:          "Sr. Cloud Architect",
				JobDescription:    "Sr. Solutions Architect for the United States Air Force (USAF). Migrated on-premise applications to AWS Govcloud, AWS Secret, and AWS Top Secret regions. Recommended, designed, and implemented cloud and application architectural enhancements to improve high availability and fault tolerance. Automated client workflows. Supported mission-critical Air Force Distributed Common Ground Systems (AF-DCGS).",
			},
			{
				CompanyName:       "Research Management Consultants, Inc.",
				CompanyCity:       "Macon",
				CompanyState:      "GA",
				CompanyStartMonth: "September",
				CompanyStartYear:  2017,
				CompanyEndMonth:   "July",
				CompanyEndYear:    2019,
				CurrentlyEmployed: false,
				JobTitle:          "Sr. Oracle DBA",
				JobDescription:    "Sr. Oracle DBA and Systems Engineer responsible for building a remote disaster recovery solution. Utilized Oracle ASM and Oracle Dataguard to replicate data in realtime to remote site. Migrated Solaris systems to Red Hat Enterprise Linux. Participated in DBA and SA on-call rotation. Received performance accolades for exceeding Key Performance Indicators (KPI) such as system uptime.",
			},
			{
				CompanyName:       "Northrop Grumman",
				CompanyCity:       "Warner Robins",
				CompanyState:      "GA",
				CompanyStartMonth: "June",
				CompanyStartYear:  2012,
				CompanyEndMonth:   "September",
				CompanyEndYear:    2017,
				CurrentlyEmployed: false,
				JobTitle:          "Programmer Analyst 3",
				JobDescription:    "Promoted to Sr. DBA. Continued prior responsibilities and lead a small team consisting of a junior DBA and Systems Administrator. Developed Amazon Web Services (AWS) solution for United States Air Force (USAF) customer. Directed and participated in full rebuild of classified systems as a result of a catastrophic failure. Mentored colleagues. Migrated Oracle 11g databases to Oracle 12c with zero downtime. Migrated Oracle Linux servers to Red Hat Enterprise Linux with zero downtime. Designed, built, deployed, and optimized a clustered enterprise search solution utilizing Apache Solr and Apache Zookeeper serving over 4,000 users worldwide.",
			},
			{
				CompanyName:       "Northrop Grumman",
				CompanyCity:       "Warner Robins",
				CompanyState:      "GA",
				CompanyStartMonth: "October",
				CompanyStartYear:  2010,
				CompanyEndMonth:   "June",
				CompanyEndYear:    2012,
				CurrentlyEmployed: false,
				JobTitle:          "Systems Engineer 2",
				JobDescription:    "Administered classified and unclassified database infrastructure, consisting of primary and standby Oracle 10g and 11g instances, as well as Microsoft SQL Server 2005, 2008, and 2008R2 servers including a remote warm site. Ensured redundancy through installing and utilizing Oracle Data Guard. Secured databases in accordance with Department of Defense (DoD) regulations.",
			},
			{
				CompanyName:       "Fred's, Inc.",
				CompanyCity:       "Dublin",
				CompanyState:      "GA",
				CompanyStartMonth: "February",
				CompanyStartYear:  2003,
				CompanyEndMonth:   "October",
				CompanyEndYear:    2010,
				CurrentlyEmployed: false,
				JobTitle:          "MIS Manager",
				JobDescription:    "IT Manager for newly-built distribution center supporting over 300 stores in the Southeast region. Lead Oracle DBA and HP-UX administrator, acting as the lone support engineer on call 7x24x365. Managed departmental budget of over $750,000 per year. Supervised hourly employees across two shifts tasked with various Tier 1 support tasks. Maintained 99.95% system uptime.",
			},
		},
		Education: []ResumeEducation{
			{
				SchoolName:         "St. Leo University",
				SchoolCity:         "St. Leo",
				SchoolState:        "FL",
				SchoolStartYear:    2007,
				SchoolEndYear:      2008,
				Major:              "MBA",
				CurrentlyAttending: false,
			},
			{
				SchoolName:         "Georgia Southern University",
				SchoolCity:         "Statesboro",
				SchoolState:        "GA",
				SchoolStartYear:    2000,
				SchoolEndYear:      2002,
				DegreeType:         "BBA",
				Major:              "Business Administration",
				MajorFocus:         "Management Information Systems",
				CurrentlyAttending: false,
			},
			{
				SchoolName:         "Middle Georgia State University",
				SchoolCity:         "Cochran",
				SchoolState:        "GA",
				SchoolStartYear:    1998,
				SchoolEndYear:      2000,
				DegreeType:         "Assoc",
				Major:              "Business Administration",
				CurrentlyAttending: false,
			},
		},
	}

	out, err := os.Create("templates/main.html")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	t, err := template.ParseGlob("templates/*")
	if err != nil {
		fmt.Println("Error on parse")
		fmt.Println(err)
	}

	err = t.Execute(out, data)
	if err != nil {
		fmt.Println("Error on execute")
		fmt.Println(err)
	}
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl.Execute(w, data)
		})
		http.ListenAndServe(":8080", nil)
	*/
}
