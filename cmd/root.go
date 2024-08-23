/*
Copyright © 2023 Bellotobiloba01@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   `Java (jarkata) Project Manager`,
	Short: "Create Java Projects",
	Long:  `Create Java Projects`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
