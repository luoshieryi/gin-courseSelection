package cache

import (
	"fmt"
	"testing"
)

/*
 @Author: as
 @Date: Creat in 21:13 2021/8/11
 @Description: magipoke-intergral
*/

func TestName(t *testing.T) {
	SetRedis()
	fmt.Println(RedisClient.Set("one","1"))
	get, _ := RedisClient.Get("one")
	fmt.Println(string(get.([]byte)))
}
