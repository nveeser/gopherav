package gopherav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

type OpenOptions struct {
	dict map[string]string
}

type AvFormat struct {
	ptr *C.struct_AVFormatContext
}

func (f *AvFormat) pointer() *C.struct_AVFormatContext {
	return (*C.struct_AVFormatContext)(unsafe.Pointer(f.ptr))
}
func (f *AvFormat) pointerRef() **C.struct_AVFormatContext {
	return (**C.struct_AVFormatContext)(unsafe.Pointer(&f.ptr))
}

func OpenInput(filename string, o *OpenOptions) (*AvFormat, error) {
	if o == nil {
		o = &OpenOptions{}
	}
	format := &AvFormat{}

	fmtPtr := (*C.struct_AVInputFormat)(C.NULL)

	cFile := C.CString(filename)
	defer C.free(unsafe.Pointer(cFile))

	dict, err := NewDictionary(o.dict)
	if err != nil {
		return nil, fmt.Errorf("error parsing optional dictionary: %w", err)
	}
	dictPtr := (**C.struct_AVDictionary)(unsafe.Pointer(&dict.ptr))

	err = fromCode(C.avformat_open_input(format.pointerRef(), cFile, fmtPtr, dictPtr))
	if err != nil {
		return nil, fmt.Errorf("error opening input: %w", err)
	}

	return format, nil
}

func OpenOutput(filename string) (*AvFormat, error) {
	format := &AvFormat{}

	cFile := C.CString(filename)
	defer C.free(unsafe.Pointer(cFile))

	err := fromCode(C.avformat_alloc_output_context2(format.pointerRef(), nil, nil, cFile))
	if err != nil {
		return nil, fmt.Errorf("error opening input: %w", err)
	}

	return format, nil
}

func (f *AvFormat) InitStreamInfo(options map[string]string) error {
	ctxPtr := (*C.struct_AVFormatContext)(unsafe.Pointer(f.ptr))

	dict, err := NewDictionary(options)
	if err != nil {
		return fmt.Errorf("error parsing optional dictionary: %w", err)
	}

	err = fromCode(C.avformat_find_stream_info(ctxPtr, dict.pointerRef()))
	if err != nil {
		return fmt.Errorf("error opening input: %w", err)
	}
	return nil
}

func (f *AvFormat) NewStream(c *Codec) *AvStream {
	s := C.avformat_new_stream(f.pointer(), c.pointer())
	if s == nil {
		return nil
	}
	return &AvStream{
		format: f,
		ptr:    s,
	}
}

func (f *AvFormat) Streams() []*AvStream {
	var cstream []*C.struct_AVStream
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&cstream)))
	slice.Cap = int(f.ptr.nb_streams)
	slice.Len = int(f.ptr.nb_streams)
	slice.Data = uintptr(unsafe.Pointer(f.ptr.streams))

	streams := make([]*AvStream, len(cstream))
	for i, cs := range cstream {
		streams[i] = &AvStream{
			format: f,
			ptr:    cs,
		}
	}
	return streams
}

func (f *AvFormat) Close() {
	C.avformat_free_context(f.ptr)
}
