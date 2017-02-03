package volume

import "strings"

type BootRecord struct {
	volumeHeader

	SystemID [32]byte
	ID       [32]byte
	Data     [1977]byte
}

func (r BootRecord) Id() string {
	return strings.TrimRight(string(r.ID[:]), "\x00")
}

func (r BootRecord) SystemId() string {
	return strings.TrimRight(string(r.SystemID[:]), "\x00")
}
