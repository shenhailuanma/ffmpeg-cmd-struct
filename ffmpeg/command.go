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

	if videoParams.Logo != nil {
		// check
		if videoParams.Logo.Source.Local != "" {
			logoCmd := fmt.Sprintf("movie=%s", videoParams.Logo.Source.Local)

			var w = 0 // logo scale size
			var h = 0 // logo scale size
			if videoParams.Logo.W != nil {
				w = *videoParams.Logo.W
			}
			if videoParams.Logo.H != nil {
				h = *videoParams.Logo.H
			}

			if w > 0 || h > 0 {
				if w > 0 {
					if h > 0 {
						logoCmd = fmt.Sprintf("%s,scale=%d:%d", logoCmd, w, h)
					} else {
						logoCmd = fmt.Sprintf("%s,scale=%d:-4", logoCmd, w)
					}
				} else {
					logoCmd = fmt.Sprintf("%s,scale=-4:%d", logoCmd, h)
				}
			}

			logoX := fmt.Sprintf("%d", videoParams.Logo.X)
			logoY := fmt.Sprintf("%d", videoParams.Logo.Y)

			if videoParams.Logo.Preset != nil {
				switch *videoParams.Logo.Preset {
				case FFmpegVideoLogoPresetLeftTop:
					logoX = "floor(0.05*main_w)"
					logoY = "floor(0.05*main_h)"
				case FFmpegVideoLogoPresetRightTop:
					logoX = "floor(main_w-overlay_w-0.05*main_w)"
					logoY = "floor(0.05*main_h)"
				case FFmpegVideoLogoPresetLeftBottom:
					logoX = "floor(0.05*main_w)"
					logoY = "floor(main_h-overlay_h-0.05*main_h)"
				case FFmpegVideoLogoPresetRightBottom:
					logoX = "floor(main_w-overlay_w-0.05*main_w)"
					logoY = "floor(main_h-overlay_h-0.05*main_h)"
				case FFmpegVideoLogoPresetLeftUpDown:
					logoX = "floor(0.05*main_w)"
					logoY = "(main_h-overlay_h)*abs(sin(t/100))"
				case FFmpegVideoLogoPresetRightUpDown:
					logoX = "floor(main_w-overlay_w-0.05*main_w)"
					logoY = "(main_h-overlay_h)*abs(sin(t*0.01))"
				case FFmpegVideoLogoPresetTopLeftRight:
					logoX = "(main_w-overlay_w)*abs(sin(t*0.01))"
					logoY = "floor(0.05*main_h)"
				case FFmpegVideoLogoPresetLeftTopRightBottom:
					logoX = "(main_w-overlay_w)*abs(sin(2*PI*t/100))"
					logoY = "(main_h-overlay_h)*abs(sin(2*PI*t/100))"
				case FFmpegVideoLogoPresetCircle:
					logoX = "(main_w-overlay_w)/2+(main_w-overlay_w)/2*sin(2*PI*t/100)"
					logoY = "(main_h-overlay_h)/2-(main_h-overlay_h)/2*cos(2*PI*t/100)"
				}
			}

			if logoX != "" && logoY != "" {
				logoCmd = fmt.Sprintf("%s [logo];[in][logo]overlay=%s:%s", logoCmd, logoX, logoY)
			}

			videoFilerList = append(videoFilerList, logoCmd)
		}
	}

	if videoParams.Delogo != nil {
		// check
		if videoParams.Delogo.X >= 0 && videoParams.Delogo.Y >= 0 && videoParams.Delogo.W >= 0 && videoParams.Delogo.H >= 0 {
			videoFilerList = append(videoFilerList, fmt.Sprintf("delogo=x=%d:y=%d:w=%d:h=%d",
				videoParams.Delogo.X, videoParams.Delogo.Y, videoParams.Delogo.W, videoParams.Delogo.H))
		}
	}

	if videoParams.Width > 0 || videoParams.Height > 0 {
		if videoParams.Width > 0 {
			if videoParams.Height > 0 {
				videoFilerList = append(videoFilerList, fmt.Sprintf("scale=%d:%d", videoParams.Width, videoParams.Height))
			} else {
				videoFilerList = append(videoFilerList, fmt.Sprintf("scale=%d:-4", videoParams.Width))
			}
		} else {
			videoFilerList = append(videoFilerList, fmt.Sprintf("scale=-4:%d", videoParams.Height))
		}
	}

	fmt.Println("videoFilerList:", videoFilerList)
	var output = ""
	if len(videoFilerList) > 0 {
		output = "-vf '"
		for filterIndex, filterOne := range videoFilerList {
			if filterIndex > 0 {
				output = output + ","
			}
			output = output + filterOne
		}
		output = output + "'"
	}

	return output
}
