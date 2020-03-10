package ogit

import (
	"log"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// CloneRepo - Clone Git Repo
func CloneRepo() billy.Filesystem {
	url := "https://github.com/danzim/manifest-test.git"

	fs := memfs.New()
	storer := memory.NewStorage()

	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		log.Fatal(err)
	}
	return fs
}
