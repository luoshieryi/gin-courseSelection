package cache

/*
 @Author: as
 @Date: Creat in 14:23 2022/2/14
 @Description: user.go
*/

type Cache interface {
	Get(k string)(interface{},error)
	Set(args ...interface{})error
}

func NewCache()Cache{
	c:=NewClient()
	if c==nil{
		return NewGoCache()
	}
	return c
}