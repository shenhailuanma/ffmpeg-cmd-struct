package ffmpeg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type FFprobeInfo struct {
	Format  FFprobeFormatInfo   `json:"format"`
	Streams []FFprobeStreamInfo `json:"streams"`
	Error   FFprobeErrorInfo    `json:"error,omitempty"`
}

type FFprobeFormatInfo struct {
	Filename   string `json:"filename"`
	NbStreams  int    `json:"nb_streams"`
	FormatName string `json:"format_name"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	ProbeScore int    `json:"probe_score"`
}

type FFprobeStreamInfo struct {
	Index              int    `json:"index"`
	CodecName          string `json:"codec_name"`
	CodecType          string `json:"codec_type"`
	BitRate            string `json:"bit_rate,omitempty"`             // video/audio
	Width              int    `json:"width,omitempty"`                // video
	Height             int    `json:"height,omitempty"`               // video
	SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`  // video
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"` // video
	PixFmt             string `json:"pix_fmt,omitempty"`              // video
	AvgFrameRate       string `json:"avg_frame_rate,omitempty"`       // video
	SampleRate         string `json:"sample_rate,omitempty"`          // audio
	Channels           int    `json:"channels,omitempty"`             // audio
}

type FFprobeErrorInfo struct {
	Code   int    `json:"code"`
	String string `json:"string"`
}

func (p *FFprobeInfo) JsonString() string {
	data, err := json.Marshal(p)
	if err == nil {
		return string(data)
	}

	return ""
}

func ProbeWithTimeout(filePath string, timeout int64) (FFprobeInfo, error) {
	var output = FFprobeInfo{}

	// ffprobe -v quiet  -print_format json -show_format -show_streams

	args := []string{"-v", "quiet", "-print_format", "json", "-show_format", "-show_streams"}
	args = append(args, filePath)

	ctx := context.Background()
	if timeout > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
		defer cancel()
	}

	fmt.Println("ProbeWithTimeout, cmd:", args)
	cmd := exec.CommandContext(ctx, "ffprobe", args...)
	cmd.Env = os.Environ()
	buf := bytes.NewBuffer(nil)
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = stdErrBuf

	err := cmd.Run()
	if err != nil {
		fmt.Println("ProbeWithTimeout, err msg:", string(stdErrBuf.Bytes()), ", error:", err.Error())
		return output, err
	}

	var outData = buf.Bytes()
	fmt.Println("ProbeWithTimeout, cmd:", string(outData))

	err = json.Unmarshal(outData, &output)

	return output, err
}
