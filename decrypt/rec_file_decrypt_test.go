package decrypt

import (
	"foxmail_password_recover/io"
	"testing"
)

func TestPasswordInRec0(t *testing.T) {
	file := io.ReadFile("E://Account.rec0")
	clientType := io.GetClientType(file)
	password := io.FindPassWord(file)

	t.Log(string(password))
	decoded := PasswordInRec0(clientType, string(password))
	t.Log(decoded)
}
