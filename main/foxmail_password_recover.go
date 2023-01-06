package main

import (
	"fmt"
	"foxmail_password_recover/decrypt"
	"foxmail_password_recover/io"
	"foxmail_password_recover/registry"
	"os"
	"path/filepath"
)

func main() {
	loadAll()

	fmt.Println("Press Enter to exit...")
	_, _ = fmt.Scanln()
}

func loadAll() {
	location := registry.GetStorageLocation()
	dir := filepath.Dir(location)

	base := filepath.Join(dir, "Storage")
	accounts, _ := os.ReadDir(base)
	for _, account := range accounts {
		info, err := account.Info()
		if err != nil {
			continue
		}
		solveFromAccountDir(base, info)
	}
}

func solveFromAccountDir(base string, account os.FileInfo) {
	fileName := filepath.Join(base, account.Name(), "Accounts", "Account.rec0")
	password := loadAllSingleFile(fileName)
	fmt.Printf("%-40s%s\n", account.Name(), password)
}

func loadAllSingleFile(fileName string) string {
	context := io.ReadFile(fileName)

	clientType := io.GetClientType(context)
	password := io.FindPassWord(context)

	return decrypt.PasswordInRec0(clientType, string(password))
}
