package unitelway

const (
	// Standard objects
	// 0xXXYY
	// XX request code, YY response code
	READ_INTERNAL_BIT    = 0x0030
	WRITE_INTERNAL_BIT   = 0x10FE
	FORCE_INTERNAL_BIT   = 0x1BFE
	READ_INTERNAL_WORD   = 0x0434
	WRITE_INTERNAL_WORD  = 0x14FE
	READ_INTERNAL_DWORD  = 0x4070
	WRITE_INTERNAL_DWORD = 0x46FE
	READ_CONSTANT_WORD   = 0x0535
	READ_CONSTANT_DWORD  = 0x4171
	READ_SYSTEM_BIT      = 0x0131
	WRITE_SYSTEM_BIT     = 0x11FE
	READ_SYSTEM_WORD     = 0x0636
	WRITE_SYSTEM_WORD    = 0x15FE
	READ_GRAFCET_BIT     = 0x2A5A
)
const (
	categoryCode = 7
)

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
