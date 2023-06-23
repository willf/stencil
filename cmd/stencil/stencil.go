package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/cbroglie/mustache"
	flag "github.com/spf13/pflag"
)

func readInput(filename string) (string, error) {
	if filename != "" {
		// Read from file
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	} else {
		// Read from standard input
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
}

func fillGoTemplate(input string, variables map[string]interface{}) (string, error) {
	// Create a new template
	tmpl, err := template.New("template").Parse(input)
	if err != nil {
		return "", err
	}

	// Execute the template with the given variables
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, variables)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func fileMoustacheTemplate(input string, variables map[string]interface{}) (string, error) {
	return mustache.Render(input, variables)
}

func fileColonTemplate(input string, variables map[string]interface{}) (string, error) {
	// Replace variable references with their values
	output := input
	for key, value := range variables {
		output = strings.ReplaceAll(output, ":"+key, fmt.Sprintf("%v", value))
	}
	return output, nil
}

func Execute() {
	var templateFile = flag.StringP("file", "f", "", "path to the template file")
	var variables = flag.StringP("variables", "v", "", "comma-separated list of key=value pairs")
	var templateType = flag.StringP("type", "t", "mustache", "type of template to use")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Stencil command: Convert templated text using variables\n\n")
		fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	Render(*templateFile, *variables, *templateType)
}

func Render(templateFile string, variables string, templateType string) {
	// Parse the variables into a map
	variablesMap := make(map[string]interface{})
	if variables != "" {
		pairs := strings.Split(variables, ",")
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				variablesMap[strings.TrimSpace(kv[0])] = kv[1]
			}
		}
	}

	var input string
	var err error

	// Read the template file
	input, err = readInput(templateFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result string

	switch {
	case templateType == "gotemplate" || templateType == "go":
		result, err = fillGoTemplate(input, variablesMap)
	case templateType == "mustache" || templateType == "moustache":
		result, err = fileMoustacheTemplate(input, variablesMap)
	case templateType == "colon":
		result, err = fileColonTemplate(input, variablesMap)
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

func main() {
	Execute()
}
