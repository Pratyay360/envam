/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// existingCmd represents the existing command
var existingCmd = &cobra.Command{
	Use:     "existing [var_name]",
	Aliases: []string{"e", "-e"},
	Short:   "add an existing record from rbw to the env file",
	Long:    `adds an existing record from rbw to the environment variables configuration file`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		varName := args[0]
		varFile := os.Getenv("ENVAM_VAR_FILE")
		varFolder := os.Getenv("ENVAM_VAR_FOLDER")
		if varFile == "" {
			varFile = os.Getenv("ENVMAN_VAR_FILE")
		}
		if varFile == "" {
			varFile = "$HOME/.config/vars"
		}
		if varFolder == "" {
			varFolder = "shell"
		}

		ExistingVar(varFile, varName, varFolder)
	},
}

func init() {
	rootCmd.AddCommand(existingCmd)
}

func ExistingVar(varFile string, varName string, varFolder string) {
	expandedPath := os.ExpandEnv(varFile)
	config, err := os.ReadFile(expandedPath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("Error reading file: %v\n", err)
			fmt.Printf("Does the file exists?")
			os.Exit(1)
		}
		config = []byte{}
	}
	if strings.Contains(string(config), varName) {
		fmt.Println("Variable already exists: " + varName)
		return
	}
	cta := "export " + varName + "=\"$(rbw get --folder " + varFolder + " " + varName + ")\"\n"

	dir := filepath.Dir(expandedPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(expandedPath, append(config, []byte(cta)...), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Variable added: " + varName)
}
