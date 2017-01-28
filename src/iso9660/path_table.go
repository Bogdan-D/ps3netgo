package iso9660

type pathTableEntry struct {
	// Length of Directory Identifier
	Len uint8
	// Extended Attribute Record Length
	ExtLen uint8

	// Location of Extent (LBA).
	// This is in a different format depending on whether this is the L-Table or M-Table.
	LBA_Addr [4]byte

	// Directory number of parent directory (an index in to the path table).
	// This is the field that limits the table to 65536 records.
	ParentNum uint16

	// Directory Identifier (name) in d-characters. Variable
	DirName []byte

	// Padding Field - contains a zero if the Length of Directory Identifier field is odd, not present otherwise.
	// This means that each table entry will always start on an even byte number.
	Padding interface{}
}

type directoryEntry struct {
	// Length of Directory Record.
	Len uint8
	//  Extended Attribute Record length.
	ExtLen uint8

	// Location of extent (LBA) in both-endian format.
	LBA_AddrLSB [4]byte
	LBA_Addr    uint32

	// Data length (size of extent) in both-endian format.
	DataLenLSB [4]byte
	DataLen    uint32

	// Recording date and time
	DateTime [7]byte

	FileFlags byte

	// File unit size for files recorded in interleaved mode, zero otherwise.
	UnitSize byte

	// Interleave gap size for files recorded in interleaved mode, zero otherwise.
	GapSize byte

	// Volume sequence number - the volume that this extent is recorded on, in 16 bit both-endian format.
	VolumeNumLSB [2]byte
	VolumeNum    uint16

	// Length of file identifier (file name).
	// This terminates with a ';' character followed by the file ID number in ASCII coded decimal ('1').
	FilenameLen uint8

	FileName string

	// Padding field - zero if length of file identifier is even, otherwise, this field is not present.
	// This means that a directory entry will always start on an even byte number.
	Padding interface{}

	// System Use - The remaining bytes up to the maximum record size of 255 may be used for extensions of ISO 9660.
	// The most common one is the System Use Share Protocol (SUSP) and its application, the Rock Ridge Interchange Protocol (RRIP).
	SystemUse interface{}
}
