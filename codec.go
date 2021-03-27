package gopherav

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <inttypes.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"fmt"
	"math/big"
	"unsafe"
)

type Codec C.struct_AVCodec

type CodecMode int

const (
	Encoder CodecMode = 0
	Decoder CodecMode = 1
)

func FindCodec(id CodecID, m CodecMode) (*Codec, error) {
	var codec *Codec
	switch m {
	case Encoder:
		codec = (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
	case Decoder:
		codec = (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
	}
	if codec == nil {
		return nil, fmt.Errorf("not codec for codec id: %d", id)
	}
	return codec, nil
}

func FindCodecByName(name string, m CodecMode) (*Codec, error) {
	cName := unsafe.Pointer(C.CString(name))
	defer C.free(cName)

	var codec *Codec
	switch m {
	case Encoder:
		codec = (*Codec)(C.avcodec_find_encoder_by_name((*C.char)(cName)))
	case Decoder:
		codec = (*Codec)(C.avcodec_find_decoder_by_name((*C.char)(cName)))
	}
	if codec == nil {
		return nil, fmt.Errorf("not codec for codec name: %s", name)
	}
	return codec, nil
}

func (c *Codec) String() string { return C.GoString(c.name) }

func (c *Codec) LongName() string { return C.GoString(c.long_name) }

func (c *Codec) pointer() *C.struct_AVCodec { return (*C.struct_AVCodec)(unsafe.Pointer(c)) }

type CodecParameters struct {
	MediaType MediaType
	CodecID   CodecID
	CodecTag  uint32
	ExtraData []byte

	Format             int
	BitRate            int64
	BitsPerCodedSample int
	BitsPerRawSample   int
	Profile, Level     int

	VideoParameters VideoCodecParameters
	AudioParameters AudioCodecParameters
}

type VideoCodecParameters struct {
	Width, Height               int
	SampleAspectRatio           *big.Rat
	FieldOrder                  AvFieldOrder
	ColorRange                  AvColorRange
	ColorPrimaries              AvColorPrimaries
	ColorTransferCharacteristic AvColorTransferCharacteristic
	ColorSpace                  AvColorSpace
	ChromaLocation              AvChromaLocation
	VideoDelay                  int
}

type AudioCodecParameters struct {
	ChannelLayout   uint64
	Channels        int
	SampleRate      int
	BlockAlign      int
	FrameSize       int
	InitialPadding  int
	TrailingPadding int
	SeekPreroll     int
}

func (p *CodecParameters) toStruct() (*C.struct_AVCodecParameters, func()) {
	var extraData unsafe.Pointer
	var release = func() {}
	if p.ExtraData != nil {
		extraData = C.CBytes(p.ExtraData)
		release = func() { C.free(extraData) }
	}
	return &C.struct_AVCodecParameters{
		codec_type:            int32(p.MediaType),
		codec_id:              uint32(p.CodecID),
		codec_tag:             C.uint(p.CodecTag),
		extradata:             (*C.uchar)(extraData),
		extradata_size:        C.int(len(p.ExtraData)),
		format:                C.int(p.Format),
		bit_rate:              C.long(p.BitRate),
		bits_per_coded_sample: C.int(p.BitsPerCodedSample),
		bits_per_raw_sample:   C.int(p.BitsPerRawSample),
		profile:               C.int(p.Profile),
		level:                 C.int(p.Level),

		// Video Only
		width:               C.int(p.VideoParameters.Width),
		height:              C.int(p.VideoParameters.Height),
		sample_aspect_ratio: toCRational(p.VideoParameters.SampleAspectRatio),
		field_order:         uint32(p.VideoParameters.FieldOrder),
		color_range:         uint32(p.VideoParameters.ColorRange),
		color_primaries:     uint32(p.VideoParameters.ColorPrimaries),
		color_trc:           uint32(p.VideoParameters.ColorTransferCharacteristic),
		color_space:         uint32(p.VideoParameters.ColorSpace),
		chroma_location:     uint32(p.VideoParameters.ChromaLocation),
		video_delay:         C.int(p.VideoParameters.VideoDelay),

		// Audio Only
		channel_layout:   C.uint64_t(p.AudioParameters.ChannelLayout),
		channels:         C.int(p.AudioParameters.Channels),
		sample_rate:      C.int(p.AudioParameters.SampleRate),
		block_align:      C.int(p.AudioParameters.BlockAlign),
		frame_size:       C.int(p.AudioParameters.FrameSize),
		initial_padding:  C.int(p.AudioParameters.InitialPadding),
		trailing_padding: C.int(p.AudioParameters.TrailingPadding),
		seek_preroll:     C.int(p.AudioParameters.SeekPreroll),
	}, release
}

func fromCCodecParameters(ap *C.struct_AVCodecParameters) *CodecParameters {
	return &CodecParameters{
		MediaType:          MediaType(ap.codec_type),
		CodecID:            CodecID(ap.codec_id),
		CodecTag:           uint32(ap.codec_tag),
		ExtraData:          C.GoBytes(unsafe.Pointer(ap.extradata), ap.extradata_size),
		Format:             int(ap.format),
		BitRate:            int64(ap.bit_rate),
		BitsPerCodedSample: int(ap.bits_per_coded_sample),
		BitsPerRawSample:   int(ap.bits_per_raw_sample),
		Profile:            int(ap.profile),
		Level:              int(ap.level),
		VideoParameters: VideoCodecParameters{
			Width:                       int(ap.width),
			Height:                      int(ap.height),
			SampleAspectRatio:           fromCRational(ap.sample_aspect_ratio),
			FieldOrder:                  AvFieldOrder(ap.field_order),
			ColorRange:                  AvColorRange(ap.color_range),
			ColorPrimaries:              AvColorPrimaries(ap.color_primaries),
			ColorTransferCharacteristic: AvColorTransferCharacteristic(ap.color_trc),
			ColorSpace:                  AvColorSpace(ap.color_space),
			ChromaLocation:              AvChromaLocation(ap.chroma_location),
			VideoDelay:                  int(ap.video_delay),
		},
		AudioParameters: AudioCodecParameters{
			ChannelLayout:   uint64(ap.channel_layout),
			Channels:        int(ap.channels),
			SampleRate:      int(ap.sample_rate),
			BlockAlign:      int(ap.block_align),
			FrameSize:       int(ap.frame_size),
			InitialPadding:  int(ap.initial_padding),
			TrailingPadding: int(ap.trailing_padding),
			SeekPreroll:     int(ap.seek_preroll),
		},
	}
}

type CodecContext struct {
	ptr *C.struct_AVCodecContext
}

func (c *CodecContext) Type() MediaType {
	return MediaType(c.ptr.codec_type)
}

func
NewCodecContext(params *CodecParameters, m CodecMode, options map[string]string) (*CodecContext, error) {
	codec, err := FindCodec(params.CodecID, m)
	if err != nil {
		return nil, err
	}

	ctxPtr := (*C.struct_AVCodecContext)(unsafe.Pointer(C.avcodec_alloc_context3(codec.pointer())))

	cParams, release := params.toStruct()
	defer release()

	err = fromCode(C.avcodec_parameters_to_context(ctxPtr, cParams))
	if err != nil {
		return nil, fmt.Errorf("error opening input: %w", err)
	}

	dict, err := NewDictionary(options)
	if err != nil {
		return nil, err
	}
	defer dict.free()

	err = fromCode(C.avcodec_open2(ctxPtr, codec.pointer(), dict.pointerRef()))
	if err != nil {
		return nil, fmt.Errorf("error opening input: %w", err)
	}

	return &CodecContext{
		ptr: ctxPtr,
	}, nil
	// AVStream *avs = avfc->streams[i];
	// AVCodec *avc = avcodec_find_decoder(avs->codecpar->codec_id);
	// AVCodecContext *avcc = avcodec_alloc_context3(*avc);
	// avcodec_parameters_to_context(*avcc, avs- > codecpar)
	// avcodec_open2(*avcc, *avc, NULL)
}

func (c *CodecContext) Close() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(c.ptr)))
	C.av_freep(unsafe.Pointer(c))
}
