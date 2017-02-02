package iso9660

import (
	"encoding/binary"
	"io"
	. "iso9660/volume"
)

type directoryRecord [34]byte

type iso struct {
	// reserved
	_ [32768]byte

	BootRecords []BootRecord

	PrimaryVol PrimaryRecord

	SupplVols []SupplementaryRecord
}

// Read iso header from Reader and return *iso struct
func NewIsoFromReader(r io.Reader) (i *iso, err error) {
	var descriptor VolumeDescriptor

	i = &iso{}

	for {
		err = binary.Read(r, binary.BigEndian, &descriptor)
		if err != nil || descriptor.IsTerminator() {
			break
		}

		if descriptor.IsReserved() {
			continue
		}

		switch {
		case descriptor.IsBoot():
			if boot := descriptor.ToBootRecord(); boot != nil {
				i.BootRecords = append(i.BootRecords, *boot)
			}
		case descriptor.IsPrimary():
			if primary := descriptor.ToPrimary(); primary != nil {
				i.PrimaryVol = *primary
			}

		case descriptor.IsSupplementary():
			if suppl := descriptor.ToSupplementary(); suppl != nil {
				i.SupplVols = append(i.SupplVols, *suppl)
			}
		}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return i, nil
}
