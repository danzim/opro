package ogit

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"

	helper "github.com/danzim/opro/helper"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// CreateProject -
func CreateProject(ci string) {

	var tpl = `apiVersion: v1
kind: Namespace
metadata:
  annotations:
  openshift.io/description: |
    {{.Description}}
  openshift.io/display-name: {{.DisplayName}}
  name: {{.Name}}
  spec:
finalizers:
- kubernetes`

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
			yamlTemplate := template.Must(template.New("tpl").Parse(tpl))

			yamlTemplate.Execute(&buffer, paramsProject)

			fs, r := CloneRepo()

			path := fmt.Sprintf("/%s/00_namespace/", ci)
			filepath := fmt.Sprintf("/%s/00_namespace/%s.yaml", ci, ci)
			filename := fmt.Sprintf("%s/00_namespace/%s.yaml", ci, ci)
			//fmt.Println(ci)
			//fmt.Println(path)
			err := fs.MkdirAll(path, 0755)
			if err != nil {
				log.Fatal(err)
			}
			//fs.MkdirAll("/test/bla", 0755)

			dir, _ := fs.ReadDir("/")
			for _, folder := range dir {
				fmt.Println(folder.Name())
			}
			yamlFile, err := fs.Create(filepath)
			if err != nil {
				fmt.Printf("File Creation failed")
				log.Fatal(err)
			}

			data := buffer.String()
			yamlFile.Write([]byte(data))
			fileinfo, _ := fs.Lstat(filepath)
			fmt.Println(fileinfo.IsDir())

			file, _ := fs.Open(filepath)
			buf := new(bytes.Buffer)
			buf.ReadFrom(file)
			s := buf.String()
			fmt.Printf(s)

			w, err := r.Worktree()
			if err != nil {
				fmt.Printf("Worktree return failed")
				log.Fatal(err)
			}

			status, _ := w.Status()
			for stPath := range status {
				fmt.Println(stPath)
			}

			_, err = w.Add(filename)
			if err != nil {
				fmt.Printf("Git add failed")
				log.Fatal(err)
			}

			status, _ = w.Status()
			fmt.Println(status)

			commit, err := w.Commit("Test Commit Go", &git.CommitOptions{
				Author: &object.Signature{
					Name:  "danzim",
					Email: "daniel.zimny90@gmail.com‚",
					When:  time.Now(),
				},
			})
			if err != nil {
				fmt.Printf("Commit failed")
				log.Fatal(err)
			}
			obj, err := r.CommitObject(commit)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(obj)

			token := helper.GetToken()
			err = r.Push(&git.PushOptions{
				Auth: &http.BasicAuth{
					Username: "bla",
					Password: token,
				},
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
