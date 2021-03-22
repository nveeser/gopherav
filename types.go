package gopherav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import "fmt"

type Rational struct {
	Num int
	Den int
}

func (r Rational) toStruct() C.struct_AVRational {
	var cr C.struct_AVRational
	cr.num = C.int(r.Num)
	cr.den = C.int(r.Den)
	return cr
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}

func fromStructRational(cr C.struct_AVRational) Rational {
	return Rational{
		Num: int(cr.num),
		Den: int(cr.den),
	}
}
