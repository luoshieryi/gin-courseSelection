package hash

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewHash(t *testing.T) {
	//实例化结构体
	hash := Bcrypt{
		cost: bcrypt.DefaultCost,
	}
	//模拟密码
	password := "123456"
	//加密
	bytes, err := hash.Make([]byte(password))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))

	//检查明文是否为加密密码的明文
	err = hash.Check(bytes, []byte(password))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("密码正确")
}
