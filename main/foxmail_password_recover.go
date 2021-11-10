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
	fmt.Printf("%v\t\t", account.Name())
	fileName := filepath.Join(base, account.Name(), "Accounts", "Account.rec0")
	loadAllSingleFile(fileName)
}

func loadAllSingleFile(fileName string) {
	context := io.ReadFile(fileName)

	clientType := io.GetClientType(context)
	password := io.FindPassWord(context)

	decoded := decrypt.PasswordInRec0(clientType, string(password))

	fmt.Printf("%v\n", decoded)
}
