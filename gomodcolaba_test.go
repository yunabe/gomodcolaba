package gomodcolaba

import "testing"

func TestVersion(t *testing.T) {
	got := Version()
	want := "v1.1.3"
	if want != got {
		t.Errorf("Version() = %q; want %q", got, want)
	}
}
