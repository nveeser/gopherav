package gopherav

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"

type (
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	CodecID                       C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	PixelFormat                   C.enum_AVPixelFormat
	AvSampleFormat                C.enum_AVSampleFormat
	MediaType                     C.enum_AVMediaType
)

const (
	MediaTypeUnknown    = MediaType(C.AVMEDIA_TYPE_UNKNOWN)
	MediaTypeVideo      = MediaType(C.AVMEDIA_TYPE_VIDEO)
	MediaTypeAudio      = MediaType(C.AVMEDIA_TYPE_AUDIO)
	MediaTypeData       = MediaType(C.AVMEDIA_TYPE_DATA)
	MediaTypeSubtitle   = MediaType(C.AVMEDIA_TYPE_SUBTITLE)
	MediaTypeAttachment = MediaType(C.AVMEDIA_TYPE_ATTACHMENT)
	MediaTypeNB         = MediaType(C.AVMEDIA_TYPE_NB)
)

func (t MediaType) String() string {
	return C.GoString(C.av_get_media_type_string(int32(t)))
}
