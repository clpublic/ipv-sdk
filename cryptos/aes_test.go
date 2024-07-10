package cryptos

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesEncryptCBC(t *testing.T) {
	enM, err := AesEncryptCBC([]byte("TestMessage"), []byte("qwertyuiop123456asdfghjk"))
	if err != nil {
		t.Error(err)
		return
	}
	bM := base64.StdEncoding.EncodeToString(enM)
	fmt.Println(bM)
	t.Log(bM)
}
