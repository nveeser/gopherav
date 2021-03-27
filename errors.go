package gopherav

//#cgo pkg-config: libavutil
//#include <libavutil/error.h>
//#include <stdlib.h>
//static const char *error2string(int code) { return av_err2str(code); }
import "C"
import (
	"errors"
	"fmt"
)

const (
	AvErrorEOF    = -('E' | ('O' << 8) | ('F' << 16) | (' ' << 24))
	AvErrorEAGAIN = -35
)

func fromCode(c C.int) error { return ErrorFromCode(int(c)) }

func ErrorFromCode(code int) error {
	if code >= 0 {
		return nil
	}
	return errors.New(C.GoString(C.error2string(C.int(code))))
}

func ErrorFromCodeMsg(code int, msg string) error {
	if code >= 0 {
		return nil
	}
	err := errors.New(C.GoString(C.error2string(C.int(code))))
	return fmt.Errorf("%s: %v", msg, err.Error())
}

type UnavailableOptionsErr struct {
	Keys []string
}

func (u *UnavailableOptionsErr) Error() string {
	return fmt.Sprintf("error: specified options were unavailable: %s", u.Keys)
}
