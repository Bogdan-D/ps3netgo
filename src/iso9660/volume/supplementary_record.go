package volume

type SupplementaryRecord struct {
	volumeHeader

	VolumeFlags int8

	// The name of the system that can act upon sectors 0x00-0x0F for the volume.
	SystemID [32]byte
	// Identification of this volume.
	ID [32]byte

	// unused
	_ [8]byte

	// Number of blocks in Little Endian
	BlockCountLSB [4]byte
	// Number of blocks in Big Endian
	BlockCount uint32

	EscapeSeq [32]byte

	// The size of the set in this logical volume (number of disks). Little Endian
	DiskCountLSB [2]byte
	// The size of the set in this logical volume (number of disks). Big Endian
	DiskCount uint16

	// The number of this disk in the Volume Set. Little Endian
	DiskNumLSB [2]byte
	// The number of this disk in the Volume Set. Big Endian
	DiskNum uint16

	// BlockSize in Little Endian
	BlockSizeLSB [2]byte
	// BlockSize in Big Endian
	BlockSize uint16

	// The size in bytes of the path table. Little Endian
	PathTableSizeLSB [4]byte
	// The size in bytes of the path table. Big Endian
	PathTableSize uint32

	// LBA location of the path table contains only little-endian values.
	PathTableSectorLSB [4]byte
	// LBA location of the optional path table contains only little-endian values.
	// Zero means that no optional path table exists.
	OptPathTableSectorLSB [4]byte

	// LBA location of the path table contains only big-endian values
	PathTableSector uint32
	// LBA location of the optional path table contains only big-endian values.
	// Zero means that no optional path table exists.
	OptPathTableSector uint32

	// Directory Record, which contains a single byte Directory Identifier
	RootDirectory [34]byte

	// Identifier of the volume set of which this volume is a member.
	VolumeName [128]byte
	// The volume publisher. For extended publisher information, the first byte should be 0x5F,
	// followed by the filename of a file in the root directory. If not specified, all bytes should be 0x20.
	PublisherFile [128]byte
	// The identifier of the person(s) who prepared the data for this volume.
	// For extended preparation information, the first byte should be 0x5F,
	// followed by the filename of a file in the root directory. If not specified, all bytes should be 0x20.
	PreparerFile [128]byte
	// Identifies how the data are recorded on this volume. For extended information, the first byte should be 0x5F,
	// followed by the filename of a file in the root directory. If not specified, all bytes should be 0x20.
	AppID [128]byte

	// Filename of a file in the root directory that contains copyright information for this volume set.
	// If not specified, all bytes should be 0x20.
	CopyrightFile [38]byte
	// Filename of a file in the root directory that contains abstract information for this volume set.
	// If not specified, all bytes should be 0x20.
	AbstractFile [36]byte
	// Filename of a file in the root directory that contains bibliographic information for this volume set.
	// If not specified, all bytes should be 0x20.
	BiblioFile [37]byte

	CreateDate    [17]byte
	ModDate       [17]byte
	ExpireDate    [17]byte //  The date and time after which this volume is considered to be obsolete. If not specified, then the volume is never considered to be obsolete.
	EffectiveDate [17]byte // The date and time after which the volume may be used. If not specified, the volume may be used immediately.

	StructureVersion int8 // The directory records and path table version (always 0x01).

	_ byte // Unused

	ApplicationUsed [512]byte //  Contents not defined by ISO 9660.

	_ [653]byte //  Reserved by ISO.
}
