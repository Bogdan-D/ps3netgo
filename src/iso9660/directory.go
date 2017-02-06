package iso9660

import (
	"bytes"
	"encoding/binary"
	"io"
)

type directoryRecordHeader struct {
	// Length of Directory Record.
	Len uint8
	//  Extended Attribute Record length.
	ExtLen uint8

	// Location of extent (LBA) in both-endian format.
	LBA_LocationLSB [4]byte
	LBA_Location    uint32

	// Data length (size of extent) in both-endian format.
	LBA_SizeLSB [4]byte
	LBA_Size    uint32

	// Recording date and time
	RecordDate [7]byte

	FileFlags uint8

	// File unit size for files recorded in interleaved mode, zero otherwise.
	InterleaveSize uint8
	// Interleave gap size for files recorded in interleaved mode, zero otherwise.
	InterleaveGap uint8

	// Volume sequence number - the volume that this extent is recorded on, in 16 bit both-endian format.
	VolumeNumLSB [2]byte
	VolumeNum    uint16

	// Length of file identifier (file name).
	// This terminates with a ';' character followed by the file ID number in ASCII coded decimal ('1').
	FilenameLen uint8
}

type directoryRecord struct {
	directoryRecordHeader

	Filename []byte

	_ byte

	System []byte
}

func (d *directoryRecord) Unpack(b []byte) {
	reader := bytes.NewReader(b)
	d.UnpackFromReader(reader)
}

func (d *directoryRecord) UnpackFromReader(r io.Reader) {
	binary.Read(r, binary.BigEndian, &d.directoryRecordHeader)

	if d.FilenameLen > 0 {
		d.Filename = make([]byte, d.FilenameLen)
		binary.Read(r, binary.BigEndian, d.Filename)
	}

	binary.Read(r, binary.BigEndian, d.System)
}
