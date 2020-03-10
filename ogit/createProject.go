package ogit

import (
	"bytes"
	"fmt"
)

var template = `apiVersion: v1
kind: Namespace
metadata:
  annotations:
  openshift.io/description: |
  description
  openshift.io/display-name: displayName
  name: {{.}}
  spec:
finalizers:
- kubernetes`

// CreateProject -
func CreateProject(ci string) {
	type ParamsProject struct {
		Name        string
		DisplayName string
		Description string
	}

	var buffer bytes.Buffer
	//var data io.Reader

	for _, singleProject := range Namespaces {
		if singleProject.CI == ci {
			name := singleProject.CI
			description := singleProject.Description
			displayName := singleProject.DisplayName

			paramsProject := ParamsProject{name, displayName, description}
			yamlTemplate := template.Must(template.New("template").Parse(template))

			yamlTemplate.Execute(&buffer, paramsProject)
			fmt.Println(buffer)
		}
	}
}
