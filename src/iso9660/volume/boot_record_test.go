package volume

import "testing"

var record BootRecord

func TestBootRecord_Id(t *testing.T) {
	var id = "BOOT ID"
	copy(record.ID[:], id)

	if record.Id() != id {
		t.Fail()
	}
}

func TestBootRecord_SystemId(t *testing.T) {
	var system_id = "BOOT SYSTEM ID"

	copy(record.SystemID[:], system_id)

	if record.SystemId() != system_id {
		t.Fail()
	}
}
