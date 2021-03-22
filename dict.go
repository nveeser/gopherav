package gopherav

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/dict.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type Dictionary struct {
	ptr *C.struct_AVDictionary
}

func NewDictionary(m map[string]string) (*Dictionary, error) {
	d := &Dictionary{}
	for k, v := range m {
		Ckey := C.CString(k)
		defer C.free(unsafe.Pointer(Ckey))

		Cvalue := C.CString(v)
		defer C.free(unsafe.Pointer(Cvalue))

		if errno := int(C.av_dict_set(&d.ptr, Ckey, Cvalue, 0)); errno > 0 {
			return nil, ErrorFromCode(errno)
		}
	}
	return d, nil
}

func (d *Dictionary) Size() int {
	return int(C.av_dict_count(d.ptr))
}

func (d *Dictionary) pointer() unsafe.Pointer {
	return unsafe.Pointer(d.ptr)
}
func (d *Dictionary) pointerRef() unsafe.Pointer {
	return unsafe.Pointer(&d.ptr)
}

func (d *Dictionary) free() {
	if d != nil && d.ptr != nil {
		C.av_dict_free(&d.ptr)
	}
}

func (d *Dictionary) toMap() map[string]string {
	m := map[string]string{}
	d.GetPrefix("", func(k, v string) {
		m[k] = v
	})
	return m
}

func (d *Dictionary) toUnavailableOptionsErr() (error, bool) {
	if d.Size() == 0 {
		return nil, false
	}

	var unvailable []string
	d.GetPrefix("", func(k, v string) {
		unvailable = append(unvailable, k)
	})
	return &UnavailableOptionsErr{unvailable}, true
}

func (d *Dictionary) GetPrefix(key string, f func(k, v string)) {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	var centry *C.struct_AVDictionaryEntry
	for {
		centry = C.av_dict_get(d.ptr, Ckey, centry, C.AV_DICT_IGNORE_SUFFIX)
		if centry == nil {
			break
		}
		key := C.GoString(centry.key)
		value := C.GoString(centry.value)
		f(key, value)
	}
}
