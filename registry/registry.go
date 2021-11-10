package registry

import "golang.org/x/sys/windows/registry"

const (
	FoxmailKey = `SOFTWARE\Aerofox\FoxmailPreview`
	Value      = "Executable"
)

func GetStorageLocation() string {
	key, err := registry.OpenKey(registry.CURRENT_USER, FoxmailKey, registry.QUERY_VALUE)
	if err != nil {
		panic("can not find the location foxmail installed")
	}

	defer func(key registry.Key) {
		_ = key.Close()
	}(key)

	value, _, err := key.GetStringValue(Value)
	if err != nil {
		panic("can not find the location foxmail installed")
	}

	return value
}
