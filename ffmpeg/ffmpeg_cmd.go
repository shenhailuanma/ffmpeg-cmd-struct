package ffmpeg

import (
	"errors"
	"github.com/shenhailuanma/ffmpeg-cmd-struct/ffmpeg/templates"
)

func FFmpegCommandGenerate(request FFmpegCommandLineStruct) (string, error) {
	return templates.GenerateCommand("transcode", templates.FFmpegTranscodeTemplate, request)
}

func CheckFFmpegTranscodeRequest(request FFmpegCommandLineStruct) error {

	if len(request.Inputs) == 0 {
		return errors.New("no input stream")
	}
	if len(request.Outputs) == 0 {
		return errors.New("no output stream")
	}

	return nil
}
