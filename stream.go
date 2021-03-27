package gopherav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"math/big"
	"unsafe"
)

type AvStream struct {
	format *AvFormat
	ptr    *C.struct_AVStream
}

func (s *AvStream) OpenCodecContext(m CodecMode, options map[string]string) (*CodecContext, error) {
	return NewCodecContext(s.CodecParameters(), m, options)
}

//int64_t av_stream_get_end_pts (const AvStream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *AvStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s.ptr)))
}

func (s *AvStream) CodecParameters() *CodecParameters {
	return fromCCodecParameters(s.ptr.codecpar)
}

func (s *AvStream) Metadata() map[string]string {
	d := &Dictionary{ptr: s.ptr.metadata}
	return d.toMap()
}

func (s *AvStream) GuessFrameRate() *big.Rat {
	ctxPtr := (*C.struct_AVFormatContext)(unsafe.Pointer(s.format.ptr))

	streamPtr := (*C.struct_AVStream)(unsafe.Pointer(s.ptr))

	rat := C.av_guess_frame_rate(ctxPtr, streamPtr, nil)
	return fromCRational(rat)
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

func (s *AvStream) AvgFrameRate() *big.Rat {
	return fromCRational(s.ptr.avg_frame_rate)
}

// func (avs *AvStream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.ptr.display_aspect_ratio))
// }

func (s *AvStream) RFrameRate() *big.Rat {
	return fromCRational(s.ptr.r_frame_rate)
}

func (s *AvStream) SampleAspectRatio() *big.Rat {
	return fromCRational(s.ptr.sample_aspect_ratio)
}

func (s *AvStream) TimeBase() *big.Rat {
	return fromCRational(s.ptr.time_base)
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

//Return the next frame of a stream.
func (c *AvStream) ReadFrame(pkt *Packet) error {
	fmtCtx := (*C.struct_AVFormatContext)(unsafe.Pointer(c.format.ptr))
	return fromCode(C.av_read_frame(fmtCtx, toCPacket(pkt)))
}
