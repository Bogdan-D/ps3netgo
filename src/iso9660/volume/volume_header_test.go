package volume

import "testing"

var header volumeHeader

func TestVolumeHeader_IsBoot(t *testing.T) {
	header.Type = BOOT_RECORD

	if !header.IsBoot() {
		t.Fail()
	}
}

func TestVolumeHeader_IsPrimary(t *testing.T) {
	header.Type = PRIMARY_VOLUME

	if !header.IsPrimary() {
		t.Fail()
	}
}

func TestVolumeHeader_IsSupplementary(t *testing.T) {
	header.Type = SUPPLEMENTARY_VOLUME

	if !header.IsSupplementary() {
		t.Fail()
	}
}

func TestVolumeHeader_IsTerminator(t *testing.T) {
	header.Type = SET_TERMINATOR

	if !header.IsTerminator() {
		t.Fail()
	}
}

func TestVolumeHeader_TypeName(t *testing.T) {
	header.Type = BOOT_RECORD
	if header.TypeName() != "Boot Record" {
		t.Fail()
	}

	header.Type = PRIMARY_VOLUME
	if header.TypeName() != "Primary Volume Descriptor" {
		t.Fail()
	}

	header.Type = SUPPLEMENTARY_VOLUME
	if header.TypeName() != "Supplementary Volume Descriptor" {
		t.Fail()
	}

	header.Type = VOLUME_PARTITION
	if header.TypeName() != "Volume Partition Descriptor" {
		t.Fail()
	}

	header.Type = SET_TERMINATOR
	if header.TypeName() != "Volume Descriptor Set Terminator" {
		t.Fail()
	}
}
