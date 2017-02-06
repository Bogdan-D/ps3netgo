package iso9660

import (
	"bytes"
	"encoding/binary"
	"io"
)

type pathTableHeader struct {
	// Length of Directory Identifier
	Len uint8
	// Extended Attribute Record Length
	ExtLen uint8

	// Location of Extent (LBA).
	// This is in a different format depending on whether this is the L-Table or M-Table.
	LBA_Location [4]byte

	// Directory number of parent directory (an index in to the path table).
	// This is the field that limits the table to 65536 records.
	ParentNum uint16
}

type pathTableEntry struct {
	pathTableHeader

	// Directory Identifier (name) in d-characters. Variable
	DirName []byte

	// Padding Field - contains a zero if the Length of Directory Identifier field is odd, not present otherwise.
	// This means that each table entry will always start on an even byte number.
	Padding interface{}
}

func (p *pathTableEntry) Unpack(b []byte) {
	reader := bytes.NewReader(b)
	p.UnpackFromReader(reader)
}

func (p *pathTableEntry) UnpackFromReader(r io.Reader) {
	if err := binary.Read(r, binary.BigEndian, &p.pathTableHeader); err != nil {
		panic(err)
	}

	to_read := p.Len - uint8(binary.Size(p))

	p.DirName = make([]byte, to_read)

	binary.Read(r, binary.BigEndian, p.DirName)
}
