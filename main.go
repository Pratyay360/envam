package main

import (
	"fmt"
	"os"
	"strings"
)

var varFile = os.Getenv("ENV_FILE")

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: program <variable-name>")
		os.Exit(1)
	}
	if varFile == "" {
		varFile = os.Getenv("HOME") + "/.config/vars"	
	}
	varName := os.Args[1]
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
	cta := "export " + varName + "=\"$(rbw get --folder shell " + varName + ")\"\n"
	err = os.WriteFile(expandedPath, append(config, []byte(cta)...), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Variable added: " + varName)
}
