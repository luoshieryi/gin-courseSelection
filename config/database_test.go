package config

import (
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {
	Mysql = (&mysql{}).Load("../config.ini").Init()
	fmt.Println(Mysql.Host)
	if Mysql.source == nil {
		t.Fail()
	}
}
