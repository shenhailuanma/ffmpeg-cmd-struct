package templates

const FFmpegTranscodeTemplate = `
{{- range $inputIndex, $inputOne := .Inputs}}
-i '{{$inputOne.Local}}'
	{{- if $inputOne.ClipStart }}
		-ss {{$inputOne.ClipStart}}
	{{- end }}
	{{- if $inputOne.ClipEnd }}
		-to {{$inputOne.ClipEnd}}
	{{- end }}
{{- end }}

{{- range $outputIndex, $outputOne := .Outputs}}
	{{- range $streamIndex, $streamOne := $outputOne.Streams}}
		{{- if eq "video" $streamOne.Kind}}
			{{- if $streamOne.Video.Codec}}
		 		-c:v {{$streamOne.Video.Codec}}
			{{- end }}
			{{- if $streamOne.Video.Preset}}
		 		-preset {{$streamOne.Video.Preset}}
			{{- end }}
			{{vfilters $streamOne.Video}}
			{{- if $streamOne.Video.Fps}}
		 		-r {{$streamOne.Video.Fps}}
			{{- end }}
			{{- if $streamOne.Video.Bitrate}}
		 		-b:v {{$streamOne.Video.Bitrate}}
			{{- end }}
			{{- if $streamOne.Video.CRF}}
		 		-crf {{$streamOne.Video.CRF}}
			{{- end }}
		{{- end }}
		{{- if eq "audio" $streamOne.Kind}}
			{{- if $streamOne.Audio.Codec}}
		 		-c:a {{$streamOne.Audio.Codec}}
			{{- end }}
			{{- if $streamOne.Audio.Channels}}
		 		-ac {{$streamOne.Audio.Channels}}
			{{- end }}
			{{- if $streamOne.Audio.SampleRate}}
		 		-ar {{$streamOne.Audio.SampleRate}}
			{{- end }}
			{{- if $streamOne.Audio.Bitrate}}
		 		-b:a {{$streamOne.Audio.Bitrate}}
			{{- end }}
		{{- end }}
    {{- end }}
	{{- if $outputOne.Format}}
		-f {{$outputOne.Format}}
	{{- end}}
		-y '{{$outputOne.Output.Local}}'
{{- end }}
`
