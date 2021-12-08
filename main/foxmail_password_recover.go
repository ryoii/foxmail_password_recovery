package main

import (
	"fmt"
	"foxmail_password_recover/decrypt"
	"foxmail_password_recover/io"
	"foxmail_password_recover/registry"
	"io/ioutil"
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
	accounts, _ := ioutil.ReadDir(base)
	for _, account := range accounts {
		solveFromAccountDir(base, account)
	}
}

func solveFromAccountDir(base string, account os.FileInfo) {
	fmt.Printf("%-40s", account.Name())
	fileName := filepath.Join(base, account.Name(), "Accounts", "Account.rec0")
	loadAllSingleFile(fileName)
}

func loadAllSingleFile(fileName string) {
	context := io.ReadFile(fileName)

	clientType := io.GetClientType(context)
	password := io.FindPassWord(context)

	decoded := decrypt.PasswordInRec0(clientType, string(password))

	fmt.Printf("%s\n", decoded)
}
