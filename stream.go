package gopherav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"
)

type AvStream struct {
	context *AvFormat
	ptr     *C.struct_AVStream
}

//Get side information from stream.
func (s *AvStream) AvStreamGetSideData(t AvPacketSideDataType, z int) *uint8 {
	cPrt := (*C.struct_AVStream)(s.ptr)
	return (*uint8)(C.av_stream_get_side_data(cPrt, (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

//Rational av_stream_get_r_frame_rate (const AvStream *s)
func (s *AvStream) GetRFrameRate() Rational {
	return fromStructRational(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s.ptr)))
}

//void av_stream_set_r_frame_rate (AvStream *s, Rational r)
func (s *AvStream) SetRFrameRate(r Rational) {
	cPtr := (*C.struct_AVStream)(s.ptr)
	rat := C.struct_AVRational{
		num: C.int(r.Num),
		den: C.int(r.Den),
	}
	C.av_stream_set_r_frame_rate(cPtr, rat)
}

//struct CodecParserContext * av_stream_get_parser (const AvStream *s)
//func (s *AvStream) AvStreamGetParser() *CodecParserContext {
//	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s.ptr)))
//}

// //char * av_stream_get_recommended_encoder_configuration (const AvStream *s)
// func (s *AvStream) AvStreamGetRecommendedEncoderConfiguration() string {
// 	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
// }

// //void av_stream_set_recommended_encoder_configuration (AvStream *s, char *configuration)
// func (s *AvStream) AvStreamSetRecommendedEncoderConfiguration( c string) {
// 	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
// }

//int64_t av_stream_get_end_pts (const AvStream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *AvStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s.ptr)))
}

func (s *AvStream) CodecParameters() *CodecParameters {
	return (*CodecParameters)(unsafe.Pointer(s.ptr.codecpar))
}

func (s *AvStream) Codec() *CodecContext {
	return &CodecContext{s.ptr.codec}
}

func (s *AvStream) Metadata() map[string]string {
	d := &Dictionary{ptr: s.ptr.metadata}
	return d.toMap()
}

//func (s *AvStream) IndexEntries() *AvIndexEntry {
//	return (*AvIndexEntry)(unsafe.Pointer(s.ptr.index_entries))
//}

//func (s *AvStream) AttachedPic() Packet {
//	return *fromCPacket(&s.ptr.attached_pic)
//}

//func (s *AvStream) SideData() *AvPacketSideData {
//	return (*AvPacketSideData)(unsafe.Pointer(s.ptr.side_data))
//}
//
//func (s *AvStream) ProbeData() AvProbeData {
//	return AvProbeData(s.ptr.probe_data)
//}

func (s *AvStream) AvgFrameRate() Rational {
	return fromStructRational(s.ptr.avg_frame_rate)
}

// func (avs *AvStream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.ptr.display_aspect_ratio))
// }

func (s *AvStream) RFrameRate() Rational {
	return fromStructRational(s.ptr.r_frame_rate)
}

func (s *AvStream) SampleAspectRatio() Rational {
	return fromStructRational(s.ptr.sample_aspect_ratio)
}

func (s *AvStream) TimeBase() Rational {
	return fromStructRational(s.ptr.time_base)
}

// func (avs *AvStream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.ptr.recommended_encoder_configuration)
// }

func (s *AvStream) CodecInfoNbFrames() int {
	return int(s.ptr.codec_info_nb_frames)
}

func (s *AvStream) Disposition() int {
	return int(s.ptr.disposition)
}

func (s *AvStream) EventFlags() int {
	return int(s.ptr.event_flags)
}

func (s *AvStream) Id() int {
	return int(s.ptr.id)
}

func (s *AvStream) Index() int {
	return int(s.ptr.index)
}

func (s *AvStream) Duration() int64 {
	return int64(s.ptr.duration)
}

func (s *AvStream) NbFrames() int64 {
	return int64(s.ptr.nb_frames)
}
