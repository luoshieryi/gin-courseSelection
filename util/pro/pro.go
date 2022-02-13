package pro

import (
	"fmt"
	"github.com/panjf2000/ants"
)

/*
 @Author: as
 @Date: Creat in 17:44 2022/2/13
 @Description: 对协程的复用
*/

type Tmp func()

var poolFunc *ants.PoolWithFunc

func SetPro() {
	p, err := ants.NewPoolWithFunc(20, func(i interface{}) {
		f, ok := i.(Tmp)
		if !ok {
			fmt.Println("What happened")
			return
		}
		f()
	})
	if err != nil {
		panic(err)
	}
	poolFunc = p
	go poolFunc.Running()
}

func Close() {
	_ = poolFunc.Release()
}

// AddTask 添加任务
func AddTask(f Tmp) error {
	return poolFunc.Invoke(f)
}
