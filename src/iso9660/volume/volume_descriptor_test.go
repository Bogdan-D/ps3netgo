package volume

import "testing"

var descriptor VolumeDescriptor

func TestVolumeHeader_ToBootRecord(t *testing.T) {
	descriptor.Type = BOOT_RECORD
	descriptor.Ident = [5]byte{'C', 'D', '0', '0', '1'}
	descriptor.Version = 1
	copy(descriptor.Data[:32], "BOOT SYSTEM")
	copy(descriptor.Data[32:], "BOOT CD")

	boot := descriptor.ToBootRecord()

	if !boot.IsBoot() {
		t.Fail()
	}

	if boot.Version != 1 {
		t.Fail()
	}

	if boot.Ident != [5]byte{'C', 'D', '0', '0', '1'} {
		t.Fail()
	}

	if boot.Id() != "BOOT CD" {
		t.Error("Wrong boot ID:", boot.ID)
	}

	if boot.SystemId() != "BOOT SYSTEM" {
		t.Error("Wrong boot System ID: ", boot.SystemID)
	}
}
