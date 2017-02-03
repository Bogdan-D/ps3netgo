package volume

const (
	BOOT_RECORD = iota
	PRIMARY_VOLUME
	SUPPLEMENTARY_VOLUME
	VOLUME_PARTITION

	SET_TERMINATOR = 255
)

type volumeHeader struct {
	Type    uint8
	Ident   [5]byte
	Version byte
}

func (d volumeHeader) TypeName() string {
	var descriptorTypes = map[uint8]string{
		0:   "Boot Record",
		1:   "Primary Volume Descriptor",
		2:   "Supplementary Volume Descriptor",
		3:   "Volume Partition Descriptor",
		255: "Volume Descriptor Set Terminator",
	}

	name, ok := descriptorTypes[d.Type]
	if !ok {
		return "Reserved"
	}

	return name
}

func (d volumeHeader) IsTerminator() bool {
	return d.Type == SET_TERMINATOR
}

func (d volumeHeader) IsSupplementary() bool {
	return d.Type == SUPPLEMENTARY_VOLUME
}

func (d volumeHeader) IsPrimary() bool {
	return d.Type == PRIMARY_VOLUME
}

func (d volumeHeader) IsBoot() bool {
	return d.Type == BOOT_RECORD
}
