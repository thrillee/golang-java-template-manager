package internals

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type NewProject struct {
	ProjectName  string
	Dir          string
	GroupId      string
	ArtifactId   string
	OgArtifactId string
	ProjectType  string
}

func processFile(filePath, ogArtifactId, artifactId string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(content), ogArtifactId, artifactId)
	err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func createVisitFileFunc(artifactId, ogArtifactId string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil
		}
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			err := processFile(path, ogArtifactId, artifactId)
			if err != nil {
				fmt.Printf("Error processing file %s: %v\n", path, err)
			}
		}
		return nil
	}
}

func createVisitFolderFunc(artifactId, ogArtifactId string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing dir %s: %v\n", path, err)
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() {
			if info.Name() == ogArtifactId {
				newPath := filepath.Join(filepath.Dir(path), artifactId)
				err := os.Rename(path, newPath)
				if err != nil {
					fmt.Printf("Error renaming folder: %v\n", err)
				} else {
					return filepath.SkipDir
				}
			}
		}

		return nil
	}
}

func HandleNewProject(d NewProject) error {
	projectDirName := strings.ReplaceAll(d.ProjectName, " ", "_")
	newProjectDir := filepath.Join(d.Dir, projectDirName)
	os.Mkdir(newProjectDir, os.ModePerm)

	repoURL := getRepoURL(d.ProjectType)
	if repoURL == "" {
		return fmt.Errorf("Project template not found")
	}

	err := cloneRepo(repoURL, newProjectDir)
	if err != nil {
		return err
	}

	visitFileFunc := createVisitFileFunc(d.ArtifactId, d.OgArtifactId)
	err = filepath.Walk(newProjectDir, visitFileFunc)

	visitFolderFunc := createVisitFolderFunc(d.ArtifactId, d.OgArtifactId)
	err = filepath.Walk(newProjectDir, visitFolderFunc)

	return err
}
