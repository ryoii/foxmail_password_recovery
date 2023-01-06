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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Press Enter to exit...")
		_, _ = fmt.Scanln()
	}()

	args := os.Args
	if len(args) > 1 {
		loadFromPath(args[1])
	} else {
		loadFromRegistry()
	}
}

// loadFromPath read the password from the path given with arguments.
// If the path is a file, it will be loaded as rec0 file.
// If the path is a directory, it will be solved as the Storage directory.
func loadFromPath(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !stat.IsDir() {
		password := loadPasswordFromFile(path)
		fmt.Println(password)
		return
	} else {
		solveStorage(path)
	}
}

// Get the Storage directory from registry
func loadFromRegistry() {
	location := registry.GetStorageLocation()
	dir := filepath.Dir(location)
	solveStorage(filepath.Join(dir, "Storage"))
}

func solveStorage(storagePath string) {
	accounts, _ := os.ReadDir(storagePath)
	for _, account := range accounts {
		info, err := account.Info()
		if err != nil {
			continue
		}
		solveAccountDir(storagePath, info)
	}
}

func solveAccountDir(base string, account os.FileInfo) {
	fileName := filepath.Join(base, account.Name(), "Accounts", "Account.rec0")
	password := loadPasswordFromFile(fileName)
	fmt.Printf("%-40s%s\n", account.Name(), password)
}

func loadPasswordFromFile(fileName string) string {
	context := io.ReadFile(fileName)

	clientType := io.GetClientType(context)
	password := io.FindPassWord(context)

	return decrypt.PasswordInRec0(clientType, string(password))
}
