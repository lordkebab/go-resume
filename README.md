# go-resume 

## About
go-resume is a resume builder written in Go. go-resume builds a single webpage with all your resume data from a JSON or YAML file.

## Instructions
To get started, follow these steps.

1. Clone the repo and `cd` to the directory
2. Build the project with `go build .`
3. Generate either a yaml or json template with either <br />
`./go-resume print -f json -o template.json` for json; or <br />
`./go-resume print -f yaml -o template.yaml` for yaml
4. You can also print the template to STDOUT with the `-s` flag. See `./go-resume print -h` for more info.
5. Edit the file containing the template and populate it with your data.
6. To build your resume, execute the following with either <br />
`./go-resume build -i filename_with_resume_data.json -o resume_webpage.html` for json; or <br />
`./go-resume build -i filename_with_resume_data.yaml -o resume_webpage.html` for yaml; or <br />
7. Your webpage will be generated in the same directory.

