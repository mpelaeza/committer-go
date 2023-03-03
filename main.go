package main

import(
	"os"
	"errors"
	"strings"
)

func main(){
	intro("Asistente para creación de commits")
	var commit string
	var err error
	var files  []string
	var commitTypeIndex  int
	var breakingChange bool

	// Load Changed an Staged Files
	var changedFiles, errorChangedFiles =  getChangedFiles()
	var stagedFiles, errorStagedFiles =  getStagedFiles()

	if(errorChangedFiles != nil || errorStagedFiles != nil){
		outro("Error: Comprueba que estás en un repositorio de git")
		os.Exit(1)
	}

	// Add Files
	if (len(stagedFiles) == 0 && len(changedFiles) > 0) {
		files, err= multiSelect("No tienes archivos preparados para hacer commit.\n Por favor selecciona los ficheros que quieres añadir:", changedFiles)
		isError(err, "Error: No hay archivos preparados para hacer commit")
		if(len(files) == 0) {
			err = errors.New("")
			isError(err, "Error: No hay archivos preparados para hacer commit")
		}
	}

	// Select Commit Type
	commitTypes := make([]string, len(COMMIT_TYPES))
	for i, t := range COMMIT_TYPES {
		commitTypes[i] = t.name
	}
	commitTypeIndex, err = selectInput("Selecciona el tipo de commit",commitTypes, commitTypesDescription)
	isError(err, "")
	if(len(files) == -1) {
			err = errors.New("")
			isError(err, "Error: No seleccionó ningún tipo")
	}

	commitType := COMMIT_TYPES[commitTypeIndex]

	// Set Commit Message
	commit, err = text("Introduce el mensaje del commit")
	isError(err, "")
	commit = commitType.emoji + " " + commitType.name + ": " + commit

	// Set Breaking Change
	if(commitType.release){
		breakingChange, err = confirm("¿Tiene en este commit cambios que rompen la compatibilidad anterior?", false)
		isError(err, "")
	}

	if(breakingChange){
		commit = commit +  " BREAKING_CHANGE"
	}

	// Should continue
		var shouldContinue bool
		shouldContinue, err = confirm("¿Quieres crear el commit con el mensaje" + commit + " ?", true)
		isError(err,"")

		if(!shouldContinue){
			outro("No se ha creado el commit")
			os.Exit(0)
		}

		gitAdd(files)
		gitCommit(commit)

	outro("✅ Commit creado con éxito" + commit)
}

func commitTypesDescription(value string, index int) string{
	commitType := COMMIT_TYPES[index]
	name := commitType.name

	return commitType.emoji + " " + strings.Repeat(" ", 9 - len(name) ) + "· " + commitType.description
}


func isError(err error, msg string){
	if(err != nil){
		if(len(msg) == 0){
			msg = "No fue posible realizar el commit"
		}
		// fmt.Printf("Prompt failed %v\n", err.Error())
		outro(msg)
		os.Exit(1)
	}
}