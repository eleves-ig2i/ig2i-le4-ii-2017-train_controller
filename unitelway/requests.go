package unitelway

import "fmt"

const (
	categoryCode = 7
)

var (
	ErrIncompatibleType = fmt.Errorf("incompatible type with UNI-TE")
)

//func WriteObject(segment, address byte, data []interface{}) ([]byte, byte, error) {
//	var objectType byte
//	switch data.(type) {
//	default:
//		return nil, 0, ErrIncompatibleType
//	case []bool:
//	case []int8:
//	case []byte:
//	case []uint16:
//	case []int16:
//	case []int32:
//	case []float32:
//		//math.Float32bits()
//		//math.Float32frombits()
//	}
//	request := []byte{WRITE_OBJECT >> 8, 7, segment}
//}

// bitNumber is between 0-255
func ReadBit(bitNumber byte) ([]byte, byte) {
	return []byte{READ_INTERNAL_BIT >> 8, categoryCode, bitNumber}, READ_INTERNAL_BIT % 256
}

// wordNumber is between 0-255
func ReadWord(wordNumber byte) ([]byte, byte) {
	return []byte{READ_INTERNAL_WORD >> 8, categoryCode, wordNumber}, READ_INTERNAL_WORD % 256
}

// dwordNumber is between 0-255
func ReadDWord(dwordNumber byte) ([]byte, byte) {
	return []byte{READ_INTERNAL_DWORD >> 8, categoryCode, dwordNumber}, READ_INTERNAL_DWORD % 256
}

// constantWordNumber is between 0-255
func ReadConstantWord(constantWordNumber byte) ([]byte, byte) {
	return []byte{READ_CONSTANT_WORD >> 8, categoryCode, constantWordNumber}, READ_CONSTANT_WORD % 256
}

// constantDWordNumber is between 0-255
func ReadConstantDWord(constantDWordNumber byte) ([]byte, byte) {
	return []byte{READ_CONSTANT_DWORD >> 8, categoryCode, constantDWordNumber}, READ_CONSTANT_DWORD % 256
}

// systemBitNumber is between 0-255
func ReadSystemBit(systemBitNumber byte) ([]byte, byte) {
	return []byte{READ_SYSTEM_BIT >> 8, categoryCode, systemBitNumber}, READ_SYSTEM_BIT % 256
}

// systemWordNumber is between 0-255
func ReadSystemWord(systemWordNumber byte) ([]byte, byte) {
	return []byte{READ_SYSTEM_WORD >> 8, categoryCode, systemWordNumber}, READ_SYSTEM_WORD % 256
}

// grafcetBitNumber is between 0-255
func ReadGrafcetBit(grafcetBitNumber byte) ([]byte, byte) {
	return []byte{READ_GRAFCET_BIT >> 8, categoryCode, grafcetBitNumber}, READ_GRAFCET_BIT % 256
}
