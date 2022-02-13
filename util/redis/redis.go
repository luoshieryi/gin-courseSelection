package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

/*
 @Author: as
 @Date: Creat in 18:52 2021/8/11
 @Description: 对redis的操作。
*/

const (
	Addr        = "localhost:6379"
	IdLeTimeout = 5
	MaxIdle     = 20
	MaxActive   = 8
)

type OptionPool struct {
	addr        string
	idLeTimeout int
	maxIdle     int
	maxActive   int
}

type PoolExt interface {
	apply(*OptionPool)
}

type tempFunc func(pool *OptionPool)

type FuncPoolExt struct {
	f tempFunc
}

func NewFuncPoolExt(f tempFunc) *FuncPoolExt {
	return &FuncPoolExt{f: f}
}

func SetAddr(addr string) PoolExt {
	return NewFuncPoolExt(func(pool *OptionPool) {
		pool.addr = addr
	})
}

func (f *FuncPoolExt) apply(pool *OptionPool) {
	f.f(pool)
}

type Client struct {
	Opt OptionPool

	pool *redis.Pool
}

var DefaultOption = OptionPool{
	addr:        Addr,
	idLeTimeout: IdLeTimeout,
	maxIdle:     MaxIdle,
	maxActive:   MaxActive,
}

var RedisClient *Client

// SetRedis 设置redis
func SetRedis(opt ...PoolExt) {
	RedisClient = NewClient()
}

func NewClient(opt ...PoolExt) *Client {
	c := Client{Opt: DefaultOption}
	for _, ext := range opt {
		ext.apply(&c.Opt)
	}
	c.setPool()
	return &c
}

func (pc *Client) setPool() {
	pc.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", pc.Opt.addr)
			if err != nil || dial == nil {
				return nil, err
			}
			return dial, nil
		},
		MaxActive:   pc.Opt.maxActive,
		MaxIdle:     pc.Opt.maxIdle,
		IdleTimeout: time.Duration(pc.Opt.idLeTimeout) * time.Second,
	}

}

// SetKey 设置本项目的key与真实姓名
func SetKey(str string) string {
	return "magipoke-intergral" + str
}

func (c *Client) Set(args ...interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(k string) (interface{}, error) {
	conn := c.pool.Get()
	defer conn.Close()
	val, err := conn.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *Client) HSet(k string, f string, v interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", k, f, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) HGet(k string, f string) (interface{}, error) {
	conn := c.pool.Get()
	defer conn.Close()
	val, err := conn.Do("HGET", k, f)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *Client) HDel(k string, f string) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("HDEL", k, f)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) HGetAll(k string) ([]string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	val, err := conn.Do("HGETALL", k)
	if err != nil {
		return nil, err
	}
	vals := val.([]interface{})
	var strs []string
	for i := 1; i < len(vals); i += 2 {
		strs = append(strs, vals[i].(string))
	}
	return strs, nil
}
