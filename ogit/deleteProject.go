package ogit

import (
	"fmt"
	"log"
	"time"

	helper "github.com/danzim/opro/helper"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// DeleteProject -
func DeleteProject(ci string) {
	fs, r := CloneRepo()
	dir, _ := fs.ReadDir("/")
	for _, folder := range dir {
		fmt.Println(folder.Name())
	}
	//path := fmt.Sprintf("%s", ci)
	filepath := fmt.Sprintf("/%s/00_namespace/%s.yaml", ci, ci)
	filename := fmt.Sprintf("%s/00_namespace/%s.yaml", ci, ci)
	fs.Remove(filepath)

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
