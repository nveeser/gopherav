package gopherav

type Transcoder struct {
	Input       string
	Output      string
	VideoCodec  string
	AudioCodec  string
	PrivateKeys map[string]string

	videoCodecContext *CodecContext
	audioCodecContext *CodecContext
}


func (tc *Transcoder) Go() error {
	tc.videoCodecContext
}