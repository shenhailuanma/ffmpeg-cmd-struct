package main

import (
	"fmt"
	"github.com/shenhailuanma/ffmpeg-cmd-struct/ffmpeg"
)

func main() {
	testInputClip()

}

func stringP(v string) *string {
	return &v
}

func testInputClip() error {
	var request = ffmpeg.FFmpegCommandLineStruct{
		Inputs: []ffmpeg.FFmpegInput{
			{
				Kind:      "file",
				Local:     "time60s.mp4",
				ClipStart: stringP("00:00:05"),
				ClipEnd:   stringP("00:00:15"),
			},
		},
		Outputs: []ffmpeg.FFmpegOutputParams{
			{
				Output: ffmpeg.FFmpegOutput{
					Kind:  "file",
					Local: "output.flv",
				},
				Format: "flv",
				Streams: []ffmpeg.FFmpegStreamParams{
					{
						Kind: "video",
						Video: &ffmpeg.FFmpegVideoStreamParams{
							Codec:  "h264",
							Preset: "slow",
							Width:  1000,
							Delogo: &ffmpeg.FFmpegVideoStreamParamDelogo{
								X: 100,
								Y: 200,
								W: 400,
								H: 400,
							},
						},
					},
					{
						Kind: "audio",
						Audio: &ffmpeg.FFmpegAudioStreamParams{
							Codec: "aac",
						},
					},
				},
			},
		},
	}

	// generate ffmpeg cmd
	cmd, err := ffmpeg.FFmpegCommandGenerate(request)
	if err != nil {
		fmt.Println("testInputClip, FFmpegCommandGenerate error:", err.Error())
		return nil
	}

	fmt.Printf("testInputClip, cmd:%s\n", cmd)

	return nil
}
