package volume

import (
	"bytes"
	"encoding/binary"
	"io"
	"unsafe"
)

// Basic type for Volume descriptor
type VolumeDescriptor struct {
	volumeHeader

	Data [2041]byte
}

// Convert volume descriptor struct to Supplementary struct
// Because all structs has same size, it's safe to use unsafe here
func (d *VolumeDescriptor) ToSupplementary() *SupplementaryRecord {
	if d.IsSupplementary() {
		var pointer = unsafe.Pointer(&d)
		return (*SupplementaryRecord)(pointer)
	}

	return nil
}

// Convert volume descriptor struct to primaryVolume.
// Because all structs has same size, it's safe to use unsafe here
func (d *VolumeDescriptor) ToPrimary() *PrimaryRecord {
	if d.IsPrimary() {
		var pointer = unsafe.Pointer(d)
		return (*PrimaryRecord)(pointer)
	}

	return nil
}

// Convert volume descriptor struct to bootRecord.
// Because all structs has same size, it's safe to use unsafe here
func (d *VolumeDescriptor) ToBootRecord() *BootRecord {
	if d.IsBoot() {
		var pointer = unsafe.Pointer(d)
		return (*BootRecord)(pointer)
	}

	return nil
}

func (d VolumeDescriptor) Pack() []byte {
	var buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, d)
	return buf.Bytes()
}

func (p *VolumeDescriptor) Unpack(buf []byte) {
	binary.Read(bytes.NewReader(buf), binary.BigEndian, p)
	return
}

func (p *VolumeDescriptor) UnpackFromReader(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, p)
}
