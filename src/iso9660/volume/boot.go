package volume

type BootRecord struct {
	volumeDescriptorHeader

	SysId [32]byte
	Id    [32]byte
	Data  [1977]byte
}
