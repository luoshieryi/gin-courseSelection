package cache

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

/*
 @Author: as
 @Date: Creat in 14:23 2022/2/14
 @Description: 缓存的使用
*/

var (
	NotFound=errors.New("key not find")
	ParamError=errors.New("param error")
)

type GoCache struct {
	c *cache.Cache
}

func NewGoCache(t ...time.Duration) *GoCache{
	var c *cache.Cache
	if len(t) != 2 {
		c = cache.New(time.Hour*24*7, time.Hour*24*8)
	} else {
		c = cache.New(t[0], t[1])
	}
	return &GoCache{c: c}
}

func (c *GoCache)Get(key string)(interface{},error){
	val, b := c.c.Get(key)
	if !b{
		return nil,NotFound
	}
	return val, nil
}

func (c *GoCache)Set(args ...interface{}) error{
	if len(args)<2{
		return ParamError
	}
	k:=args[0].(string)
	if len(args)==2{
		c.c.Set(k,args[1],cache.NoExpiration)
	}else {
		c.c.Set(k,args[1],args[2].(time.Duration))
	}
	return nil
}