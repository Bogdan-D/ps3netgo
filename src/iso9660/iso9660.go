package iso9660

import (
	"encoding/binary"
	"fmt"
	"io"
)

var descriptorTypes = map[int]string{
	0:   "Boot Record",
	1:   "Primary Volume Descriptor",
	2:   "Supplementary Volume Descriptor",
	3:   "Volume Partition Descriptor",
	255: "Volume Descriptor Set Terminator",
}

type descriptorHeader struct {
	Type    int8
	Ident   [5]byte
	Version byte
}

func (d descriptorHeader) TypeName() string {
	var i = int(d.Type)

	name, ok := descriptorTypes[i]
	if !ok {
		return "Reserved"
	}

	return name
}

type directoryRecord [34]byte

type bootRecord struct {
	descriptorHeader

	SysId [32]byte
	Id    [32]byte
	Data  [1977]byte
}

type basicRecord struct {
	descriptorHeader

	Data [2041]byte
}

type isoHeader struct {
	System     [32768]byte
	BootRecord bootRecord
	PrimaryVol primaryVolume
}

type iso struct {
	isoHeader
	Records []basicRecord
}

// Read header from Reader and return *iso struct
func NewIsoFromReader(r io.Reader) (*iso, error) {
	var i = &iso{}

	err := binary.Read(r, binary.BigEndian, &i.isoHeader)
	if err != nil {
		return nil, err
	}

	var record = basicRecord{}

	for err = binary.Read(r, binary.BigEndian, &record); err == nil; {
		fmt.Print(record.TypeName())

		if record.TypeName() == "Reserved" {
			continue
		}

		i.Records = append(i.Records, record)

		if record.TypeName() == "Volume Descriptor Set Terminator" {
			break
		}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return i, nil
}
