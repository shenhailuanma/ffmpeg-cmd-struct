package ffmpeg

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
