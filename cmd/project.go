package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thrillee/java-project-templates/internals"
)

var simpleWebCMD = &cobra.Command{
	Use:   "simple-web",
	Short: "Create A Simple Jakrata Web Project",
	Long:  `Create A Simple Jakrata Web Project`,
	Run:   handleSimpleWebCMD,
}

func init() {
	rootCmd.AddCommand(simpleWebCMD)

	simpleWebCMD.Flags().StringP("project", "p", "", "Project Name")
	simpleWebCMD.Flags().StringP("groupId", "g", "", "Group ID")
	simpleWebCMD.Flags().StringP("artifactId", "a", "", "Artifact ID")
	simpleWebCMD.Flags().StringP("dir", "d", "", "Directory to create project")

	simpleWebCMD.MarkFlagRequired("groupId")
	simpleWebCMD.MarkFlagRequired("artifactId")
	simpleWebCMD.MarkFlagRequired("dir")
	simpleWebCMD.MarkFlagRequired("project")
}

func handleSimpleWebCMD(simpleWebCMD *cobra.Command, args []string) {
	projectName, err := simpleWebCMD.Flags().GetString("project")
	if err != nil {
		log.Fatal(err)
	}

	groupId, err := simpleWebCMD.Flags().GetString("groupId")
	if err != nil {
		log.Fatal(err)
	}

	artifactId, err := simpleWebCMD.Flags().GetString("artifactId")
	if err != nil {
		log.Fatal(err)
	}

	dir, err := simpleWebCMD.Flags().GetString("dir")
	if err != nil {
		log.Fatal(err)
	}

	data := internals.NewProject{
		ProjectType:  "simple-web",
		OgArtifactId: "simple_web",
		ArtifactId:   artifactId,
		GroupId:      groupId,
		Dir:          dir,
		ProjectName:  projectName,
	}

	internals.HandleNewProject(data)
}
