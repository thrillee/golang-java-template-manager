package internals

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func cloneRepo(repoURL, destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Repository cloned successfully to %s\n", destination)
	return nil
}
