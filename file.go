package goff

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type OpenOptions struct {
	dict map[string]string
}

type AvFormat struct {
	ptr *C.struct_AVFormatContext
}

func Open(filename string, o *OpenOptions) (*AvFormat, error) {
	if o == nil {
		o = &OpenOptions{}
	}
	var ctx *C.struct_AVFormatContext
	ctxPtr := (**C.struct_AVFormatContext)(unsafe.Pointer(&ctx))

	fmtPtr := (*C.struct_AVInputFormat)(C.NULL)

	cFile := C.CString(filename)
	defer C.free(unsafe.Pointer(cFile))

	dict, err := NewDictionary(o.dict)
	if err != nil {
		return nil, fmt.Errorf("error parsing optional dictionary: %w", err)
	}
	dictPtr := (**C.struct_AVDictionary)(unsafe.Pointer(&dict.ptr))

	errval := int(C.avformat_open_input(ctxPtr, cFile, fmtPtr, dictPtr))
	err = ErrorFromCode(errval)
	if err != nil {
		return nil, fmt.Errorf("error opening input: %w", err)
	}

	return &AvFormat{ptr: ctx}, nil
}

func (f *AvFormat) InitStreamInfo(options map[string]string) error {
	ctxPtr := (*C.struct_AVFormatContext)(unsafe.Pointer(f.ptr))

	dict, err := NewDictionary(options)
	if err != nil {
		return fmt.Errorf("error parsing optional dictionary: %w", err)
	}
	dictPtr := (**C.struct_AVDictionary)(unsafe.Pointer(&dict.ptr))
	errval := int(C.avformat_find_stream_info(ctxPtr, dictPtr))
	err = ErrorFromCode(errval)
	if err != nil {
		return fmt.Errorf("error opening input: %w", err)
	}
	return nil
}

func (f *AvFormat) Close() {
	C.avformat_free_context(f.ptr)
}
