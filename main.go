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
	texteditor string
	gopath = os.Getenv("GOPATH") //getting value of $GOPATH from environment variable
)

func main(){
	
	if len(os.Args) < 2 {
		fmt.Printf("Create new go project and open in text editor.\nUsage:\n\t%v\t-p [Project_name] -t [Text_Editor] --path [Project_path]\n",os.Args[0])
		os.Exit(1)
	}
	
	//parsing flags
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
		}else{
			projectPath = *path + *projectName
		}
	}

	var texteditorCmd string
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
	
	//checking the project name specified by user is already exists or not
	//if exists then print the error and exit 
	_,err := os.Stat(projectPath)
	if err == nil {
		err := fmt.Errorf("Same Project name exist %v.Please choose another name",*projectName)
		fmt.Println("Error: "+err.Error())
		os.Exit(1)
	}
	
	//making project dir specified by user in -path flag and 
	//opening it in text editor specified by user in -t flag
	_ = os.MkdirAll(projectPath,os.ModePerm)
	fmt.Printf("succesfully create project %v at %v\n",*projectName,projectPath)
	cmd := exec.Command(texteditorCmd,projectPath)
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
}
