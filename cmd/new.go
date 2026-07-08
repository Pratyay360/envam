/*
Copyright © 2026 Pratyay360 <pratyaymustafi@outlook.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new [var_name]",
	Aliases: []string{"n", "-n"},
	Short:   "add a new record to rbw and adds it to the env file",
	Long:    `adds a record to rbw and references that through the env vars`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		varName := args[0]
		varFile := os.Getenv("ENVAM_VAR_FILE")
		if varFile == "" {
			varFile = os.Getenv("ENVMAN_VAR_FILE")
		}
		if varFile == "" {
			varFile = "$HOME/.config/vars"
		}

		NewVar(varFile, varName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func NewVar(varFile, varName string) {
	varValue := os.Getenv(varName)
	varFolder := os.Getenv("ENVAM_VAR_FOLDER")
	if varValue == "" {
		fmt.Println("var not found:", varName)
		if varFolder == "" {
			varFolder = "shell"
		}
		cmd := exec.Command("rbw", "add", "--folder", varFolder, varName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to execute rbw command: %v\n", err)
			os.Exit(1)
		}
	}
	ExistingVar(varFile, varName, varFolder)
}
