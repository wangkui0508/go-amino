package amino

import (
	"io"
	aminoOrig "github.com/tendermint/go-amino/go-amino"
)

//type CodecIfc interface {
//	MarshalBinaryBare(o interface{}) ([]byte, error)
//	MarshalBinaryLengthPrefixed(o interface{}) ([]byte, error)
//	MarshalBinaryLengthPrefixedWriter(w io.Writer, o interface{}) (n int64, err error)
//	MarshalJSON(o interface{}) ([]byte, error)
//	MarshalJSONIndent(o interface{}, prefix, indent string) ([]byte, error)
//	MustMarshalBinaryBare(o interface{}) []byte
//	MustMarshalBinaryLengthPrefixed(o interface{}) []byte
//	MustMarshalJSON(o interface{}) []byte
//	MustUnmarshalBinaryBare(bz []byte, ptr interface{})
//	MustUnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{})
//	MustUnmarshalJSON(bz []byte, ptr interface{})
//	PrintTypes(out io.Writer) error
//	RegisterConcrete(o interface{}, name string, copts *ConcreteOptions)
//	RegisterInterface(ptr interface{}, iopts *InterfaceOptions)
//	Seal() *Codec
//	UnmarshalBinaryBare(bz []byte, ptr interface{}) error
//	UnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) error
//	UnmarshalBinaryLengthPrefixedReader(r io.Reader, ptr interface{}, maxSize int64) (n int64, err error)
//	UnmarshalJSON(bz []byte, ptr interface{}) error
//}

type Codec = aminoOrig.Codec

const Typ3_ByteLength = aminoOrig.Typ3_ByteLength

func NewCodec() *Codec {
	return aminoOrig.NewCodec()
}

func DeepCopy(o interface{}) (r interface{}) {
	r = aminoOrig.DeepCopy(o)
	return
}

func EncodeByteSlice(w io.Writer, bz []byte) (err error) {
	err = aminoOrig.EncodeByteSlice(w, bz)
	return
}

func EncodeUvarint(w io.Writer, u uint64) (err error) {
	err = aminoOrig.EncodeUvarint(w, u)
	return
}

func MarshalBinaryBare(o interface{}) ([]byte, error) {
	return aminoOrig.MarshalBinaryBare(o)
}

func MustMarshalBinaryLengthPrefixed(o interface{}) []byte {
	return aminoOrig.MustMarshalBinaryLengthPrefixed(o)
}

func UnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) error {
	return aminoOrig.UnmarshalBinaryLengthPrefixed(bz, ptr)
}

func UvarintSize(u uint64) int {
	return aminoOrig.UvarintSize(u)
}


