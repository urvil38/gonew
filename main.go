package main

import (
	"strings"
	"os/exec"
	"flag"
	"fmt"
	"os"
)

var (
	projectName string
	projectPath string
	gopath = os.Getenv("GOPATH")
	texteditor string
	texteditorCmd string
)

func main(){
	
	if len(os.Args) < 2 {
		fmt.Printf("Create new go project and open in text editor.\nUsage:\n\t%v\t-p [Project_name] -t [Text_Editor] --path [Project_path]\n",os.Args[0])
		os.Exit(1)
	}
	
	projectName := flag.String("p","","specify new project name")
	projectPath = gopath + "/src/" + *projectName
	path := flag.String("path",projectPath,"specify path of the project")
	texteditor := flag.String("t","vscode","specify text editor to open project default is vscode.\n\toption:\n\t\tsublime , vscode , atom")
	flag.Parse()
	
	if *projectName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	if *path != "" {
		if !strings.HasSuffix(*path,"/") {
			projectPath = *path + "/" + *projectName
		}
	}

	switch (*texteditor) {
	case "atom":
		texteditorCmd = "atom"
		break
	case "sublime":
		texteditorCmd = "subl"
		break
	default:
		texteditorCmd = "code"
	}
	
	
	_,err := os.Stat(projectPath)
	if err == nil {
		err := fmt.Errorf("Same Project name exist %v.Please choose another name",*projectName)
		fmt.Println("Error: "+err.Error())
		os.Exit(1)
	}
	
	_ = os.MkdirAll(projectPath,os.ModePerm)
	fmt.Printf("succesfully create project %v at %v\n",*projectName,projectPath)
	cmd := exec.Command(texteditorCmd,projectPath)
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
}
