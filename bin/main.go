package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"internet-cafe/pkg/command"
)

const (
	ps1 = "$ "
)

func main() {
	// read if file input exist
	if len(os.Args) > 1 && os.Args[1] != "" {
		cmdFile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer cmdFile.Close()
		cmdScanner := bufio.NewScanner(cmdFile)
		for cmdScanner.Scan() {
			cmdInput := cmdScanner.Text()
			cmdInput = strings.TrimRight(cmdInput, "\n")
			fmt.Print(ps1)
			fmt.Println(cmdInput)
			if cmdInput != "" {
				executeCommand(cmdInput)
			}
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(ps1)
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, "\n")
		if cmdInput != "" {
			executeCommand(cmdInput)
		}
	}
}

func executeCommand(input string) {
	mainCmd, err := command.InitCommand(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = mainCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
