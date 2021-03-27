package gopherav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"math/big"
	"unsafe"
)

type Packet C.struct_AVPacket

func (p *Packet) unref() {
	C.av_package_unref(toCPacket(p))
}

func toCPacket(pkt *Packet) *C.struct_AVPacket {
	return (*C.struct_AVPacket)(unsafe.Pointer(pkt))
}

func fromCPacket(pkt *C.struct_AVPacket) *Packet {
	return (*Packet)(unsafe.Pointer(pkt))
}

func toCRational(r *big.Rat) C.struct_AVRational {
	if !r.Num().IsInt64() || !r.Denom().IsInt64() {
		panic("only int64 int supported")
	}
	n, d := r.Num().Int64(), r.Denom().Int64()
	if n == 0 && d == 1 {
		return C.struct_AVRational{
			num: 0,
			den: 0,
		}
	}
	return C.struct_AVRational{
		num: C.int(n),
		den: C.int(d),
	}
}

func fromCRational(cr C.struct_AVRational) *big.Rat {
	if cr.den == 0 {
		return &big.Rat{}
	}
	return big.NewRat(int64(cr.num), int64(cr.den))
}
