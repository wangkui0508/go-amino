package amino

import (
	"io"
	aminoOrig "github.com/tendermint/go-amino/go-amino"
)

type CodecIfc interface {
	MarshalBinaryBare(o interface{}) ([]byte, error)
	MarshalBinaryLengthPrefixed(o interface{}) ([]byte, error)
	MarshalBinaryLengthPrefixedWriter(w io.Writer, o interface{}) (n int64, err error)
	MarshalJSON(o interface{}) ([]byte, error)
	MarshalJSONIndent(o interface{}, prefix, indent string) ([]byte, error)
	MustMarshalBinaryBare(o interface{}) []byte
	MustMarshalBinaryLengthPrefixed(o interface{}) []byte
	MustMarshalJSON(o interface{}) []byte
	MustUnmarshalBinaryBare(bz []byte, ptr interface{})
	MustUnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{})
	MustUnmarshalJSON(bz []byte, ptr interface{})
	PrintTypes(out io.Writer) error
	RegisterConcrete(o interface{}, name string, copts *aminoOrig.ConcreteOptions)
	RegisterInterface(ptr interface{}, iopts *aminoOrig.InterfaceOptions)
	Seal() *aminoOrig.Codec
	UnmarshalBinaryBare(bz []byte, ptr interface{}) error
	UnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) error
	UnmarshalBinaryLengthPrefixedReader(r io.Reader, ptr interface{}, maxSize int64) (n int64, err error)
	UnmarshalJSON(bz []byte, ptr interface{}) error
}

type Codec struct {
	impl CodecIfc
}

func (cdc *Codec) MarshalBinaryBare(o interface{}) ([]byte, error) {
	return cdc.impl.MarshalBinaryBare(o)
}
func (cdc *Codec) MarshalBinaryLengthPrefixed(o interface{}) ([]byte, error) {
	return cdc.impl.MarshalBinaryLengthPrefixed(o)
}
func (cdc *Codec) MarshalBinaryLengthPrefixedWriter(w io.Writer, o interface{}) (n int64, err error) {
	return cdc.impl.MarshalBinaryLengthPrefixedWriter(w, o)
}
func (cdc *Codec) MarshalJSON(o interface{}) ([]byte, error) {
	return cdc.impl.MarshalJSON(o)
}
func (cdc *Codec) MarshalJSONIndent(o interface{}, prefix, indent string) ([]byte, error) {
	return cdc.impl.MarshalJSONIndent(o, prefix, indent)
}
func (cdc *Codec) MustMarshalBinaryBare(o interface{}) []byte {
	return cdc.impl.MustMarshalBinaryBare(o)
}
func (cdc *Codec) MustMarshalBinaryLengthPrefixed(o interface{}) []byte {
	return cdc.impl.MustMarshalBinaryLengthPrefixed(o)
}
func (cdc *Codec) MustMarshalJSON(o interface{}) []byte {
	return cdc.impl.MustMarshalJSON(o)
}
func (cdc *Codec) MustUnmarshalBinaryBare(bz []byte, ptr interface{}) {
	cdc.impl.MustUnmarshalBinaryBare(bz, ptr)
}
func (cdc *Codec) MustUnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) {
	cdc.impl.MustUnmarshalBinaryLengthPrefixed(bz, ptr)
}
func (cdc *Codec) MustUnmarshalJSON(bz []byte, ptr interface{}) {
	cdc.impl.MustUnmarshalJSON(bz, ptr)
}
func (cdc *Codec) PrintTypes(out io.Writer) error {
	return cdc.impl.PrintTypes(out)
}
func (cdc *Codec) RegisterConcrete(o interface{}, name string, copts *aminoOrig.ConcreteOptions) {
	cdc.impl.RegisterConcrete(o, name, copts)
}
func (cdc *Codec) RegisterInterface(ptr interface{}, iopts *aminoOrig.InterfaceOptions) {
	cdc.impl.RegisterInterface(ptr, iopts)
}
func (cdc *Codec) Seal() *Codec {
	cdc.impl = cdc.impl.Seal()
	return cdc
}
func (cdc *Codec) UnmarshalBinaryBare(bz []byte, ptr interface{}) error {
	return cdc.impl.UnmarshalBinaryBare(bz, ptr)
}
func (cdc *Codec) UnmarshalBinaryLengthPrefixed(bz []byte, ptr interface{}) error {
	return cdc.impl.UnmarshalBinaryLengthPrefixed(bz, ptr)
}
func (cdc *Codec) UnmarshalBinaryLengthPrefixedReader(r io.Reader, ptr interface{}, maxSize int64) (n int64, err error) {
	return cdc.impl.UnmarshalBinaryLengthPrefixedReader(r, ptr, maxSize)
}
func (cdc *Codec) UnmarshalJSON(bz []byte, ptr interface{}) error {
	return cdc.impl.UnmarshalJSON(bz, ptr)
}

const Typ3_ByteLength = aminoOrig.Typ3_ByteLength

func NewCodec() *Codec {
	return &Codec{impl:aminoOrig.NewCodec()}
}

func DeepCopy(o interface{}) (r interface{}) {
	r = aminoOrig.DeepCopy(o)
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


func EncodeByteSlice(w io.Writer, bz []byte) (err error) {
	err = aminoOrig.EncodeByteSlice(w, bz)
	return
}

func EncodeUvarint(w io.Writer, u uint64) (err error) {
	err = aminoOrig.EncodeUvarint(w, u)
	return
}

func ByteSliceSize(bz []byte) int {
	return aminoOrig.ByteSliceSize(bz)
}

func EncodeInt8(w io.Writer, i int8) (err error) {
	err = aminoOrig.EncodeInt8(w, i)
	return
}
func EncodeInt16(w io.Writer, i int16) (err error) {
	err = aminoOrig.EncodeInt16(w, i)
	return
}
func EncodeInt32(w io.Writer, i int32) (err error) {
	err = aminoOrig.EncodeInt32(w, i)
	return
}
func EncodeInt64(w io.Writer, i int64) (err error) {
	err = aminoOrig.EncodeInt64(w, i)
	return
}
func EncodeVarint(w io.Writer, i int64) (err error) {
	err = aminoOrig.EncodeVarint(w, i)
	return
}
func EncodeByte(w io.Writer, b byte) (err error) {
	err = aminoOrig.EncodeByte(w, b)
	return
}
func EncodeUint8(w io.Writer, u uint8) (err error) {
	err = aminoOrig.EncodeUint8(w, u)
	return
}
func EncodeUint16(w io.Writer, u uint16) (err error) {
	err = aminoOrig.EncodeUint16(w, u)
	return
}
func EncodeUint32(w io.Writer, u uint32) (err error) {
	err = aminoOrig.EncodeUint32(w, u)
	return
}
func EncodeUint64(w io.Writer, u uint64) (err error) {
	err = aminoOrig.EncodeUint64(w, u)
	return
}
func EncodeBool(w io.Writer, b bool) (err error) {
	err = aminoOrig.EncodeBool(w, b)
	return
}
func EncodeFloat32(w io.Writer, f float32) (err error) {
	err = aminoOrig.EncodeFloat32(w, f)
	return
}
func EncodeFloat64(w io.Writer, f float64) (err error) {
	err = aminoOrig.EncodeFloat64(w, f)
	return
}
func EncodeString(w io.Writer, s string) (err error) {
	err = aminoOrig.EncodeString(w, s)
	return
}
func DecodeInt8(bz []byte) (i int8, n int, err error) {
	i, n, err = aminoOrig.DecodeInt8(bz)
	return
}
func DecodeInt16(bz []byte) (i int16, n int, err error) {
	i, n, err = aminoOrig.DecodeInt16(bz)
	return
}
func DecodeInt32(bz []byte) (i int32, n int, err error) {
	i, n, err = aminoOrig.DecodeInt32(bz)
	return
}
func DecodeInt64(bz []byte) (i int64, n int, err error) {
	i, n, err = aminoOrig.DecodeInt64(bz)
	return
}
func DecodeVarint(bz []byte) (i int64, n int, err error) {
	i, n, err = aminoOrig.DecodeVarint(bz)
	return
}
func DecodeByte(bz []byte) (b byte, n int, err error) {
	b, n, err = aminoOrig.DecodeByte(bz)
	return
}
func DecodeUint8(bz []byte) (u uint8, n int, err error) {
	u, n, err = aminoOrig.DecodeUint8(bz)
	return
}
func DecodeUint16(bz []byte) (u uint16, n int, err error) {
	u, n, err = aminoOrig.DecodeUint16(bz)
	return
}
func DecodeUint32(bz []byte) (u uint32, n int, err error) {
	u, n, err = aminoOrig.DecodeUint32(bz)
	return
}
func DecodeUint64(bz []byte) (u uint64, n int, err error) {
	u, n, err = aminoOrig.DecodeUint64(bz)
	return
}
func DecodeUvarint(bz []byte) (u uint64, n int, err error) {
	u, n, err = aminoOrig.DecodeUvarint(bz)
	return
}
func DecodeBool(bz []byte) (b bool, n int, err error) {
	b, n, err = aminoOrig.DecodeBool(bz)
	return
}
func DecodeFloat32(bz []byte) (f float32, n int, err error) {
	f, n, err = aminoOrig.DecodeFloat32(bz)
	return
}
func DecodeFloat64(bz []byte) (f float64, n int, err error) {
	f, n, err = aminoOrig.DecodeFloat64(bz)
	return
}
func DecodeByteSlice(bz []byte) (bz2 []byte, n int, err error) {
	bz2, n, err = aminoOrig.DecodeByteSlice(bz)
	return
}
func DecodeString(bz []byte) (s string, n int, err error) {
	s, n, err = aminoOrig.DecodeString(bz)
	return
}
func VarintSize(i int64) int {
	return aminoOrig.VarintSize(i)
}
