package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/cbroglie/mustache"
)

type TemplateType int

const (
	GoTemplate TemplateType = iota
	MustacheTemplate
	ColonTemplate
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

func fillGoTemplate(input string, variables map[string]string) (string, error) {
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

func fileMoustacheTemplate(input string, variables map[string]string) (string, error) {
	return mustache.Render(input, variables)
}

func fileColonTemplate(input string, variables map[string]string) (string, error) {
	// Replace variable references with their values
	output := input
	for key, value := range variables {
		output = strings.ReplaceAll(output, ":"+key, fmt.Sprintf("%v", value))
	}
	return output, nil
}

func Usage() {
	fmt.Println("Usage: stencil [OPTIONS]")
	fmt.Println("Stencil command: Convert templated text using variables")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -f --file <file>   path to a template file (default: stdin)")
	fmt.Println("  -g --go            use Go template syntax")
	fmt.Println("  -m --mustache      use Mustache template syntax (default)")
	fmt.Println("  -c --colon         use colon template syntax")
	fmt.Println("  -h --help          print this help message")
	fmt.Println("Other flags are passed as key=value pairs for use in the template")
}

func Execute() {
	var templateFile string
	var variables map[string]string = make(map[string]string)
	var useGoTemplate bool
	var useMustacheTemplate bool
	var useColonTemplate bool
	var helpme bool

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "-f", "--file":
			i++
			if i >= len(os.Args) {
				fmt.Println("Missing template file")
				os.Exit(1)
			}
			templateFile = os.Args[i]
		case "-g", "--go", "--gotemplate":
			useGoTemplate = true
		case "-m", "--mustache":
			useMustacheTemplate = true
		case "-c", "--colon":
			useColonTemplate = true
		case "-h", "--help":
			helpme = true
		default:
			// Handle key-value pairs
			if strings.Contains(arg, "=") {
				parts := strings.SplitN(arg, "=", 2)
				key := strings.TrimLeft(parts[0], "-")
				value := parts[1]
				variables[key] = value
			} else {
				key := strings.TrimLeft(arg, "-")
				i++
				if i >= len(os.Args) {
					fmt.Println("Missing value for key", key)
					os.Exit(1)
				}
				value := os.Args[i]
				variables[key] = value
			}
		}
	}

	if helpme {
		Usage()
		os.Exit(0)
	}

	var templateType TemplateType
	if useGoTemplate {
		templateType = GoTemplate
	} else if useMustacheTemplate {
		templateType = MustacheTemplate
	} else if useColonTemplate {
		templateType = ColonTemplate
	} else {
		templateType = MustacheTemplate
	}

	Render(templateFile, variables, templateType)
}

func Render(templateFile string, variables map[string]string, templateType TemplateType) {

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
	case templateType == GoTemplate:
		result, err = fillGoTemplate(input, variables)
	case templateType == MustacheTemplate:
		result, err = fileMoustacheTemplate(input, variables)
	case templateType == ColonTemplate:
		result, err = fileColonTemplate(input, variables)
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
