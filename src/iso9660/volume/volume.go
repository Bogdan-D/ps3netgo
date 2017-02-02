package volume

import "unsafe"

var descriptorTypes = map[uint8]string{
	0:   "Boot Record",
	1:   "Primary Volume Descriptor",
	2:   "Supplementary Volume Descriptor",
	3:   "Volume Partition Descriptor",
	255: "Volume Descriptor Set Terminator",
}

type volumeDescriptorHeader struct {
	Type    uint8
	Ident   [5]byte
	Version byte
}

func (d volumeDescriptorHeader) TypeName() string {
	name, ok := descriptorTypes[d.Type]
	if !ok {
		return "Reserved"
	}

	return name
}

func (d volumeDescriptorHeader) IsTerminator() bool {
	return d.TypeName() == "Volume Descriptor Set Terminator"
}

func (d volumeDescriptorHeader) IsReserved() bool {
	return d.TypeName() == "Reserved"
}

func (d volumeDescriptorHeader) IsSupplementary() bool {
	return d.TypeName() == "Supplementary Volume Descriptor"
}

func (d volumeDescriptorHeader) IsPrimary() bool {
	return d.TypeName() == "Primary Volume Descriptor"
}

func (d volumeDescriptorHeader) IsBoot() bool {
	return d.TypeName() == "Boot Record"
}

// Convert volume descriptor struct to Supplementary struct
// Because all of them has same size, it's safe to use unsafe here
func (d volumeDescriptorHeader) ToSupplementary() *SupplementaryRecord {
	if d.IsSupplementary() {
		var pointer = unsafe.Pointer(&d)
		var volumePointer = (*SupplementaryRecord)(pointer)

		return volumePointer
	}

	return nil
}

// Convert volume descriptor struct to primaryVolume.
// Because all of descriptors has same size, it's safe to use unsafe here
func (d volumeDescriptorHeader) ToPrimary() *PrimaryRecord {
	if d.IsPrimary() {
		var pointer = unsafe.Pointer(&d)
		var volumePointer = (*PrimaryRecord)(pointer)

		return volumePointer
	}

	return nil
}

// Convert volume descriptor struct to bootRecord.
// Because all of descriptors has same size, it's safe to use unsafe here
func (d volumeDescriptorHeader) ToBootRecord() *BootRecord {
	if d.IsBoot() {
		var pointer = unsafe.Pointer(&d)
		var volumePointer = (*BootRecord)(pointer)

		return volumePointer
	}

	return nil
}

type VolumeDescriptor struct {
	volumeDescriptorHeader

	Data [2041]byte
}
