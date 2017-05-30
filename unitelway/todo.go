package unitelway

const (
	// 0xXXYY
	// XX request code, YY response code
	// Usage général
	IDENTIFICATION   = 0x0F3F
	READ_CPU         = 0x4F7F
	PROTOCOL_VERSION = 0x3060
	MIRROR           = 0xFAFB



	// Module d’E/S
	READ_DIGITAL_MODULE_IMAGE  = 0x4979
	WRITE_DIGITAL_MODULE_IMAGE = 0x4A7A
	READ_STATUS_MODULE         = 0x4474
	READ_IO_CHANNEL            = 0x4373
	WRITE_IO_CHANNEL           = 0x4878



	// Modes de marche
	RUN  = 0x24FE
	STOP = 0x25FE
	INIT = 0x3363

	OPEN_DOWNLOAD  = 0x3A6A
	WRITE_DOWNLOAD = 0x3B6B
	CLOSE_DOWNLOAD = 0x3C6C
	OPEN_UPLOAD    = 0x3D6D
	READ_UPLOAD    = 0x3E6E
	CLOSE_UPLOAD   = 0x3F6F
	BACKUP         = 0x4575

	// Sémaphores
	RESERVE    = 0x1DFE
	RELEASE    = 0x1EFE
	I_AM_ALIVE = 0x2DFE
)
