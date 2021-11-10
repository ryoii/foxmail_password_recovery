package registry

import "testing"

func TestGetStorageLocation(t *testing.T) {
	location := GetStorageLocation()
	t.Log(location)
}
