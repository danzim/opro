package ogit

import (
	"fmt"
	"log"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// CloneRepo - Clone Git Repo
func CloneRepo() (billy.Filesystem, *git.Repository) {
	url := "https://github.com/danzim/manifest-test.git"

	fs := memfs.New()
	storer := memory.NewStorage()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		fmt.Printf("Cloning of Repo failed")
		log.Fatal(err)
	}
	return fs, r
}
