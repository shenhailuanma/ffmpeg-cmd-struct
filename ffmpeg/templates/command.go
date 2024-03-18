package templates

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func GenerateCommand(name string, templateText string, data interface{}) (string, error) {

	tmpl := template.New(name).Funcs(template.FuncMap{})

	tmpl, err := tmpl.Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	var cmdString = buf.String()

	// split by '\n'
	cmdStringList := strings.Split(cmdString, "\n")
	var cmdStringOutput = ""
	for _, cmdStringOne := range cmdStringList {
		// trim space and connect all strings
		cmdNewOne := strings.TrimSpace(cmdStringOne)
		if len(cmdNewOne) > 0 {
			fmt.Printf("cmdNewOne:%s\n", cmdNewOne)
			cmdStringOutput = cmdStringOutput + " " + cmdNewOne
		}
	}

	return cmdStringOutput, nil
}
