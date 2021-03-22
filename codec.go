package gopherav

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <inttypes.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"reflect"
	"unsafe"
)

type Codec C.struct_AVCodec

func FindDecoderCodec(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

func FindEncoderCodec(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

func (c *Codec) String() string {
	return C.GoString(c.name)
}

func (c *Codec) LongName() string {
	return C.GoString(c.long_name)
}

type CodecParameters C.struct_AVCodecParameters

func (cp *CodecParameters) CodecId() CodecId {
	return *((*CodecId)(unsafe.Pointer(&cp.codec_id)))
}

func (cp *CodecParameters) MediaType() MediaType {
	return *((*MediaType)(unsafe.Pointer(&cp.codec_type)))
}

func (cp *CodecParameters) Width() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.width))))
}

func (cp *CodecParameters) Height() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.height))))
}

func (cp *CodecParameters) Channels() int {
	return *((*int)(unsafe.Pointer(&cp.channels)))
}

func (cp *CodecParameters) SampleRate() int {
	return *((*int)(unsafe.Pointer(&cp.sample_rate)))
}

type CodecContext struct {
	ptr *C.struct_AVCodecContext
}

func (c *CodecContext) Type() MediaType {
	return MediaType(c.ptr.codec_type)
}

func OpenCodec(*Codec) *CodecContext {
	return nil
}

func (c *CodecContext) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(c.ptr)))
	C.av_freep(unsafe.Pointer(c))
}

func (c *CodecContext) SetBitRate(br int64) {
	c.ptr.bit_rate = C.int64_t(br)
}

func (c *CodecContext) GetCodecId() CodecId {
	return CodecId(c.ptr.codec_id)
}

func (c *CodecContext) SetCodecId(codecId CodecId) {
	c.ptr.codec_id = C.enum_AVCodecID(codecId)
}

func (c *CodecContext) GetCodecType() MediaType {
	return MediaType(c.ptr.codec_type)
}

func (c *CodecContext) SetCodecType(ctype MediaType) {
	c.ptr.codec_type = C.enum_AVMediaType(ctype)
}

func (c *CodecContext) GetTimeBase() Rational {
	return fromStructRational(c.ptr.time_base)
}

func (c *CodecContext) SetTimeBase(timeBase Rational) {
	c.ptr.time_base.num = C.int(timeBase.Num)
	c.ptr.time_base.den = C.int(timeBase.Den)
}

func (c *CodecContext) GetWidth() int {
	return int(c.ptr.width)
}

func (c *CodecContext) SetWidth(w int) {
	c.ptr.width = C.int(w)
}

func (c *CodecContext) GetHeight() int {
	return int(c.ptr.height)
}

func (c *CodecContext) SetHeight(h int) {
	c.ptr.height = C.int(h)
}

func (c *CodecContext) GetPixelFormat() PixelFormat {
	return PixelFormat(C.int(c.ptr.pix_fmt))
}

func (c *CodecContext) SetPixelFormat(fmt PixelFormat) {
	c.ptr.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

func (c *CodecContext) GetFlags() int {
	return int(c.ptr.flags)
}

func (c *CodecContext) SetFlags(flags int) {
	c.ptr.flags = C.int(flags)
}

func (c *CodecContext) GetMeRange() int {
	return int(c.ptr.me_range)
}

func (c *CodecContext) SetMeRange(r int) {
	c.ptr.me_range = C.int(r)
}

func (c *CodecContext) GetMaxQDiff() int {
	return int(c.ptr.max_qdiff)
}

func (c *CodecContext) SetMaxQDiff(v int) {
	c.ptr.max_qdiff = C.int(v)
}

func (c *CodecContext) GetQMin() int {
	return int(c.ptr.qmin)
}

func (c *CodecContext) SetQMin(v int) {
	c.ptr.qmin = C.int(v)
}

func (c *CodecContext) GetQMax() int {
	return int(c.ptr.qmax)
}

func (c *CodecContext) SetQMax(v int) {
	c.ptr.qmax = C.int(v)
}

func (c *CodecContext) GetQCompress() float32 {
	return float32(c.ptr.qcompress)
}

func (c *CodecContext) SetQCompress(v float32) {
	c.ptr.qcompress = C.float(v)
}

func (c *CodecContext) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.ptr.extradata)),
		Len:  int(c.ptr.extradata_size),
		Cap:  int(c.ptr.extradata_size),
	}
	return *((*[]byte)(unsafe.Pointer(&header)))
}

func (c *CodecContext) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	c.ptr.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	c.ptr.extradata_size = C.int(header.Len)
}
