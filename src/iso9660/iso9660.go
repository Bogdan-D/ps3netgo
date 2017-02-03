package iso9660

import (
	"io"
	. "iso9660/volume"
)

type iso struct {
	// reserved
	System [32768]byte

	BootRecords []BootRecord

	PrimaryVol PrimaryRecord

	SupplVols []SupplementaryRecord
}

// Read iso header from Reader and return *iso struct
func NewIsoFromReader(r io.Reader) (i *iso, err error) {
	i = &iso{}

	if _, err = r.Read(i.System[:]); err != nil {
		return nil, err
	}

	var descriptor VolumeDescriptor

loop:
	for {
		if err = descriptor.UnpackFromReader(r); err != nil {
			break
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
		case descriptor.IsTerminator():
			break loop
		}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return i, nil
}
