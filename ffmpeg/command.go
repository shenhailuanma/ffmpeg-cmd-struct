package ffmpeg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

func GenerateCommand(name string, templateText string, data interface{}) (string, error) {

	tmpl := template.New(name).Funcs(template.FuncMap{
		"vfilters": formatVideoFilters,
	})

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
			//fmt.Printf("cmdNewOne:%s\n", cmdNewOne)
			cmdStringOutput = cmdStringOutput + " " + cmdNewOne
		}
	}

	return cmdStringOutput, nil
}

func formatVideoFilters(input interface{}) string {
	var videoParams = FFmpegVideoStreamParams{}
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("formatVideoFilters, json Marshal error:", err.Error())
		return ""
	}

	err = json.Unmarshal(jsonData, &videoParams)
	if err != nil {
		fmt.Println("formatVideoFilters, json Unmarshal error:", err.Error())
		return ""
	}

	var videoFilerList = []string{}

	if videoParams.Delogo != nil {
		videoFilerList = append(videoFilerList, "delogo")
	}

	if videoParams.Width > 0 || videoParams.Height > 0 {
		videoFilerList = append(videoFilerList, "scale")
	}

	return ""
}
