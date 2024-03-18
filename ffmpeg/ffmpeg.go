package ffmpeg

type FFmpegGlobalParams struct {
	Overwrite bool `json:"overwrite"`
	NoStream  bool `json:"noStream"` // -vn, as an input option, blocks all video streams
}

const (
	FFmpegInputOutputTypeFILE = "file" // local file
	FFmpegInputOutputTypeHTTP = "http" // http url file
	FFmpegInputOutputTypeOSS  = "oss"
	FFmpegInputOutputTypeFTP  = "ftp"
)

type FFmpegVideoStreamParams struct {
	Map     string `json:"map"` // strams map, http://ffmpeg.org/ffmpeg-all.html#Advanced-options
	Codec   string `json:"codec"`
	Preset  string `json:"preset"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Bitrate string `json:"bitrate"`
	Fps     int    `json:"fps"`
	CRF     int    `json:"crf"`
}
type FFmpegAudioStreamParams struct {
	Map        string `json:"map"` // strams map, http://ffmpeg.org/ffmpeg-all.html#Advanced-options
	Codec      string `json:"codec"`
	Channels   int    `json:"channles"`
	SampleRate int    `json:"sample_rate"`
	Bitrate    string `json:"bitrate"`
}

type FFmpegStreamParams struct {
	Kind  string                   `json:"kind"` // video, audio, subtitle, data
	Video *FFmpegVideoStreamParams `json:"video"`
	Audio *FFmpegAudioStreamParams `json:"audio"`
}

type FFmpegOutputParams struct {
	Output  FFmpegOutput         `json:"output"`
	Format  string               `json:"format"`
	Streams []FFmpegStreamParams `json:"streams"`
}

/*
url:
  - if kind is file, eg: /tmp/folder/video.mp4
  - if kind is url, eg: https://example.com/video.mp4
  - if kind is oss, eg: https://presigned.url
*/
type FFmpegOutput struct {
	Kind  string `json:"kind"`  // kind: file(local file), http(http/https url), oss, ftp , defined in: FFmpegInputOutputType*
	URL   string `json:"url"`   // remote resource URL, when FFmpegInputOutputTypeHTTP used.
	Local string `json:"local"` // local file path,  when FFmpegInputOutputTypeFILE used. eg: /tmp/folder/video.mp4
}

type FFmpegInput struct {
	Kind      string  `json:"kind"`       // kind: file(local file), http(http/https url), oss, ftp , defined in: FFmpegInputOutputType*
	URL       string  `json:"url"`        // remote resource URL, when FFmpegInputOutputTypeHTTP used.
	Local     string  `json:"local"`      // local file path,  when FFmpegInputOutputTypeFILE used. eg: /tmp/folder/video.mp4
	ClipStart *string `json:"clip_start"` // seek the start of this input file to position. position must be a time duration specification, see (ffmpeg-utils)the Time duration section in the ffmpeg-utils(1) manual.
	ClipEnd   *string `json:"clip_end"`   // seek the end of this input file to position. position must be a time duration specification, see (ffmpeg-utils)the Time duration section in the ffmpeg-utils(1) manual.
}

type FFmpegCommandLineStruct struct {
	Inputs  []FFmpegInput        `json:"inputs"`
	Outputs []FFmpegOutputParams `json:"outputs"`
	Globals FFmpegGlobalParams   `json:"globals"`
}
