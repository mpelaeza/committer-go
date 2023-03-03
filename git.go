package main

import (
	// "fmt"
	"os/exec"
	"strings"
)

func cleanStdout(stdout string) []string {
	stdout = strings.TrimSpace(stdout) 
	var result = strings.Split(stdout, "\n")
	return result
}

func formatFiles(stdout string) []string{
	if(stdout == ""){
		return nil
	}
	var lines = cleanStdout(stdout)	
	var files = make([]string, len(lines) )
	for i := 0; i < len(lines); i++ {
		//e.g:  ?? README.md 
		aux := strings.Split(lines[i], " " )
		files[i] = 	aux[len(aux) - 1]
	}
	return files
}

func getChangedFiles() ([]string, error) {
	var cmd = exec.Command("git", "status",  "--porcelain")
	var stdout, err = cmd.Output()
	var changedFiles = formatFiles(string(stdout))
	if(err != nil){
		return nil, err
	}
	return changedFiles, nil
}

func getStagedFiles() ([]string, error) {
	var cmd = exec.Command("git", "diff",  "--cached", "--name-only")
	var stdout, err = cmd.Output()
	var stagedFiles = formatFiles(string(stdout))
	if(err != nil){
		return nil, err
	}
	return stagedFiles, nil
}

func gitAdd(files []string){
	var addCommant = "add " + strings.Join(files, " ")
	var cmd = exec.Command("git", strings.Split(addCommant, " ")...)
	cmd.Output()
	// var stdout, err = cmd.Output()
	// fmt.Println(string(stdout))
	// fmt.Println(err)
	// fmt.Printf("Prompt failed1 %v\n", err.Error())
}

func gitCommit(commit string){
	var cmd = exec.Command("git", "commit", "-m", commit)
	cmd.Output()
	// var stdout, err = cmd.Output()
	// fmt.Println(string(stdout))
	// fmt.Printf("Prompt failed2 %v\n", err.Error())
}