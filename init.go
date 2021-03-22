package goff

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <libavformat/avformat.h>
import "C"

func init() {
	C.av_register_all()
}
