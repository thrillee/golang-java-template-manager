package internals

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func cloneRepo(r *repo, destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:           r.url,
		Progress:      os.Stdout,
		SingleBranch:  true,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", r.branch)),
	})
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Repository cloned successfully to %s\n", destination)
	return nil
}
