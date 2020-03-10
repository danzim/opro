package ogit

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// Namespace - exported namespace struct
type Namespace struct {
	CI          string `json:"CI"`
	DisplayName string `json:"DisplayName"`
	Description string `json:"Description"`
}

// AllNamespaces - exported namespaces
type AllNamespaces []Namespace

var namespaces = AllNamespaces{}

// GetProjects - Get Projects in Git Repo
func GetProjects() AllNamespaces {
	fs := CloneRepo()
	dir, err := fs.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}
	namespaces = nil
	for _, folder := range dir {
		matched, err := regexp.MatchString(`ci-.*`, folder.Name())
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			ci := folder.Name()
			file, err := fs.Open(fmt.Sprintf("/%s/00_namespace/%s.yaml", ci, ci))
			if err != nil {
				log.Fatal(err)
			}

			buf := new(bytes.Buffer)
			buf.ReadFrom(file)
			s := buf.String()

			file.Close()

			result := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(s), &result)
			metadata := result["metadata"].(map[interface{}]interface{})

			name := fmt.Sprintln(metadata["name"])
			description := fmt.Sprintln(metadata["openshift.io/description"])
			displayName := fmt.Sprintln(metadata["openshift.io/display-name"])

			regex, err := regexp.Compile("\n\n")
			if err != nil {
				log.Fatal(err)
			}

			description = regex.ReplaceAllString(description, "\n")
			name = strings.TrimSuffix(name, "\n")
			description = strings.TrimSuffix(description, "\n")
			displayName = strings.TrimSuffix(displayName, "\n")

			var newNamespace = Namespace{
				CI:          name,
				DisplayName: displayName,
				Description: description,
			}
			namespaces = append(namespaces, newNamespace)
		}
	}
	return namespaces
}
