package nbt

import (
	"io"

	"github.com/MJKWoolnough/byteio"
)

type Decoder struct {
	r byteio.EndianReader
}

func NewDecoder(r io.Reader) Decoder {
	return NewDecoderEndian(byteio.BigEndianReader{r})
}

func NewDecoderEndian(e byteio.EndianReader) Decoder {
	return Decoder{r: e}
}

func (d Decoder) DecodeTag() (Tag, error) {
	t, _, err := d.r.ReadUint8()
	if err != nil {
		err = ReadError{"named TagId", err}
		return Tag{}, err
	}
	tagID := TagID(t)
	if tagID == TagEnd {
		return Tag{data: end{}}, nil
	}
	n, err := d.DecodeString()
	if err != nil {
		err = ReadError{"name", err}
		return Tag{}, err
	}
	data, err := d.decodeData(tagID)
	if err != nil {
		return Tag{}, err
	}
	return Tag{name: string(n), data: data}, nil
}

func (d Decoder) decodeData(tagID TagID) (Data, error) {
	var (
		data Data
		err  error
	)
	switch tagID {
	case TagByte:
		data, err = d.DecodeByte()
	case TagShort:
		data, err = d.DecodeShort()
	case TagInt:
		data, err = d.DecodeInt()
	case TagLong:
		data, err = d.DecodeLong()
	case TagFloat:
		data, err = d.DecodeFloat()
	case TagDouble:
		data, err = d.DecodeDouble()
	case TagByteArray:
		data, err = d.DecodeByteArray()
	case TagString:
		data, err = d.DecodeString()
	case TagList:
		data, err = d.DecodeList()
	case TagCompound:
		data, err = d.DecodeCompound()
	case TagIntArray:
		data, err = d.DecodeIntArray()
	default:
		err = UnknownTag{tagID}
	}
	if err != nil {
		if _, ok := err.(*ReadError); !ok {
			err = &ReadError{tagID.String(), err}
		}
		return nil, err
	}
	return data, nil
}

func (d Decoder) DecodeByte() (Byte, error) {
	b, _, err := d.r.ReadInt8()
	return Byte(b), err
}

func (d Decoder) DecodeShort() (Short, error) {
	s, _, err := d.r.ReadInt16()
	return Short(s), err
}

func (d Decoder) DecodeInt() (Int, error) {
	i, _, err := d.r.ReadInt32()
	return Int(i), err
}

func (d Decoder) DecodeLong() (Long, error) {
	l, _, err := d.r.ReadInt64()
	return Long(l), err
}

func (d Decoder) DecodeFloat() (Float, error) {
	f, _, err := d.r.ReadFloat32()
	return Float(f), err
}

func (d Decoder) DecodeDouble() (Double, error) {
	do, _, err := d.r.ReadFloat64()
	return Double(do), err
}

func (d Decoder) DecodeByteArray() (ByteArray, error) {
	l, _, err := d.r.ReadUint32()
	if err != nil {
		return nil, err
	}
	data := make([]byte, l)
	_, err = io.ReadFull(d.r, data)
	if err != nil {
		return nil, err
	}
	ba := make([]int8, l)
	for i := uint32(0); i < l; i++ {
		ba[i] = int8(data[i])
	}
	return ByteArray(ba), nil
}

func (d Decoder) DecodeString() (String, error) {
	l, _, err := d.r.ReadUint16()
	if err != nil {
		return "", err
	}
	data := make([]byte, l)
	_, err = io.ReadFull(d.r, data)
	if err != nil {
		return "", err
	}
	return String(data), nil
}

func (d Decoder) DecodeList() (*List, error) {
	t, _, err := d.r.ReadUint8()
	if err != nil {
		return nil, err
	}
	tagID := TagID(t)
	l, _, err := d.r.ReadUint32()
	if err != nil {
		return nil, err
	}
	data := make([]Data, l)
	for i := uint32(0); i < l; i++ {
		data[i], err = d.decodeData(tagID)
		if err != nil {
			return nil, err
		}
	}
	return &List{
		tagID,
		data,
	}, nil
}

func (d Decoder) DecodeCompound() (Compound, error) {
	data := make(Compound, 0)
	for {
		t, err := d.DecodeTag()
		if err != nil {
			return nil, err
		}
		if t.TagID() == TagEnd {
			break
		}
		data = append(data, t)
	}
	return data, nil
}

func (d Decoder) DecodeIntArray() (IntArray, error) {
	l, _, err := d.r.ReadUint32()
	if err != nil {
		return nil, err
	}
	ints := make(IntArray, l)
	for i := uint32(0); i < l; i++ {
		ints[i], _, err = d.r.ReadInt32()
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}