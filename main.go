package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Project struct {
	name   string
	path   string
	editor string
}

func main() {

	var p Project

	flag.StringVar(&p.name, "name", "", "project name")
	flag.StringVar(&p.name, "n", "", "project name")
	flag.StringVar(&p.editor, "e", "vscode", "specify text editor to open project. option: sublime, vscode, atom")
	flag.StringVar(&p.editor, "editor", "vscode", "specify text editor to open project. option: sublime, vscode, atom")
	flag.StringVar(&p.path, "path", "", "project path")
	flag.StringVar(&p.path, "p", "", "project path")
	flag.Parse()

	if p.name == "" {
		flag.Usage()
		os.Exit(1)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" && p.path == "" {
		fmt.Printf("Error: GOPATH environment variable is not set and no path provided\n")
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if gopath != "" {
		p.path = filepath.Join(gopath, "src", p.path, p.name)
	} else {
		p.path = filepath.Join(pwd, p.path, p.name)
	}

	var texteditorCmd string
	switch p.editor {
	case "atom":
		texteditorCmd = "atom"
	case "sublime":
		texteditorCmd = "subl"
	default:
		texteditorCmd = "code"
	}

	//checking the project name specified by user is already exists or not
	//if exists then print the error and exit
	if _, err = os.Stat(p.path); err == nil {
		fmt.Printf("Error: Same project name \"%v\" already exist. Please choose another name.\n", p.name)
		os.Exit(1)
	}

	//making project dir specified by user in -path flag and
	//opening it in text editor specified by user in -t flag
	_ = os.MkdirAll(p.path, os.ModePerm)
	fmt.Printf("Succesfully created project %v at %v\n", p.name, p.path)
	cmd := exec.Command(texteditorCmd, p.path)
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
}
