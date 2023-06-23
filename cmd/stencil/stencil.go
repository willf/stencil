package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/cbroglie/mustache"
)

func fillGoTemplate(templateFile string, variables map[string]interface{}) (string, error) {
	// Read the template file
	template, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	// Execute the template with the given variables
	var buf bytes.Buffer
	err = template.Execute(&buf, variables)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func fileMoustacheTemplate(templateFile string, variables map[string]interface{}) (string, error) {
	return mustache.RenderFile(templateFile, variables)
}

func fileColonTemplate(templateFile string, variables map[string]interface{}) (string, error) {
	// Read the template file
	templateBytes, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return "", err
	}

	// Convert the template to a string
	template := string(templateBytes)

	// Replace variable references with their values
	for key, value := range variables {
		template = strings.ReplaceAll(template, ":"+key, fmt.Sprintf("%v", value))
	}

	return template, nil
}

func main() {
	// Define flags for the template file and variables
	templateFile := flag.String("template", "", "path to the template file")
	variables := flag.String("variables", "", "comma-separated list of key=value pairs")
	templateType := flag.String("type", "mustache", "type of template to use")
	flag.Parse()

	// Parse the variables into a map
	variablesMap := make(map[string]interface{})
	if *variables != "" {
		pairs := strings.Split(*variables, ",")
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				variablesMap[strings.TrimSpace(kv[0])] = kv[1]
			}
		}
	}

	var result string
	var err error

	switch {
	case *templateType == "gotemplate" || *templateType == "go":
		result, err = fillGoTemplate(*templateFile, variablesMap)
	case templateType == nil || *templateType == "mustache" || *templateType == "moustache":
		result, err = fileMoustacheTemplate(*templateFile, variablesMap)
	case *templateType == "colon":
		result, err = fileColonTemplate(*templateFile, variablesMap)
	default:
		fmt.Println("Invalid template type")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
