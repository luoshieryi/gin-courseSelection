package cache

import (
	"fmt"
	"testing"
)

/*
 @Author: as
 @Date: Creat in 14:47 2022/2/14
 @Description: user.go
*/

func TestNewGoCache(t *testing.T) {
	c := NewGoCache()
	c.Set("1","2")
	val,err:=c.Get("1")
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(val.(string))
}