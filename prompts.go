package main

import(
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"errors"
	"strings"
)

// var templates = &promptui.PromptTemplates{
// 	Prompt: "{{ . }}",
// 	Valid:   "◆ {{ . | green }} ",
// 	Invalid: "▲ {{ . | yellow }} ",
// 	Success: "◇ {{ . | bold }} ",
// }


func intro(intro string) {
	fmt.Printf("┌ %s \n", intro)
	fmt.Println("│")
}
 
func outro(outro string) {
	fmt.Println("|")
	fmt.Printf("└ %s \n", outro)
}

func text(label string) (string, error){
	result := ""
	prompt := []*survey.Question{
		{
		Name: "",
		Prompt: &survey.Input{Message: label},
		Validate: textNonBlank,
		// Templates: templates,
		},
	}
	err := survey.Ask(prompt, &result)
	// result, err := prompt.Run()
	return result, err
}

func confirm(label string, defaultValue bool) (bool, error){
	var result  int = 0
	if(defaultValue){
		result = 1
	}
	prompt := &survey.Select{
    Message: label,
		Options: []string{"No", "Si"},
		Default: result,

	}
 err := survey.AskOne(prompt, &result)
 return result == 1, err

}

func selectInput(label string, options []string, descFunction func(string, int) string) (int, error){
	var result  int = -1
	prompt := &survey.Select{
    Message: label,
		Options: options,
		Description: descFunction,

	}
 err := survey.AskOne(prompt, &result)
 return result, err
}


func multiSelect(label string, options []string) ([]string, error){
	result := []string{}
	prompt := &survey.MultiSelect{
    Message: label,
			Options: options,
	}
 err := survey.AskOne(prompt, &result)
 return result, err
}


func textNonBlank(val interface{}) error{
	if str, ok := val.(string) ; !ok || len(strings.TrimSpace(str)) == 0 {
		return errors.New("el mensaje no puede estar vacio")
	}
	return nil
}