package unitelway

import "fmt"

const (
	// generic objects
	// 0xXXYY
	// XX request code, YY response code
	READ_GENERIC_OBJECT  = 0x82B2
	WRITE_GENERIC_OBJECT = 0x83B3
	READ_OBJECT          = 0x3666
	WRITE_OBJECT         = 0x37FE
	READ_OBJECT_LIST     = 0x3868

	// object types
	SystemByte = 0x0001

	InternalBit = 0x6405
	SystemBit   = 0x6406

	InternalByte = 0x6806
	InternalWord = 0x6807
	DoubleWord   = 0x6808
	SimpleFloat  = 0x680A
	DoubleFloat  = 0x680B

	ConstantWord       = 0x6907
	ConstantDoubleWord = 0x6908
	SystemWord         = 0x6A07
	SystemDoubleWord   = 0x6A08
)

var (
	ErrWrongObjectType  = fmt.Errorf("given object type doesn't exist")
	ErrIncompatibleType = fmt.Errorf("incompatible type with UNI-TE")
)

func WriteObject(objectType, address uint16, v interface{}) ([]byte, byte, error) {
	request := []byte{
		WRITE_OBJECT >> 8, 7,
		byte(objectType >> 8), byte(objectType % 256),
		byte(address % 256), byte(address >> 8),
	}
	expected := byte(WRITE_OBJECT % 256)

	switch v := v.(type) {
	case []bool:
		switch objectType {
		case SystemBit:
		case InternalBit:
		default:
			return nil, 0, ErrWrongObjectType
		}
	case []byte:
		switch objectType {
		case SystemByte:
		case InternalByte:
		default:
			return nil, 0, ErrWrongObjectType
		}
		request = append(request, byte(len(v)%256), byte(len(v)>>8))
		return append(request, v...), expected, nil
	case []uint16:
		switch objectType {
		case SystemWord:
		case InternalWord:
		case ConstantWord:
		default:
			return nil, 0, ErrWrongObjectType
		}
		request = append(request, byte(len(v)%256), byte(len(v)>>8))
		for elt := range v {
			request = append(request, byte(elt%256), byte(elt>>8))
		}
		return request, expected, nil
	}
	return nil, 0, ErrIncompatibleType
}

//case []int16:
//case []int32:
//case []uint32:
//	switch objectType {
//	case SystemDoubleWord:
//	case DoubleWord:
//	case ConstantDoubleWord:
//	default:
//		return nil, 0, ErrWrongObjectType
//	}
//case []float32:
//	switch objectType {
//	case SimpleFloat:
//	default:
//		return nil, 0, ErrWrongObjectType
//	}
//case []float64:
//	switch objectType {
//	case DoubleFloat:
//	default:
//		return nil, 0, ErrWrongObjectType
//	}
//math.Float32bits()
//math.Float32frombits()
