package pro

import (
	"fmt"
	"sync"
	"testing"
)

/*
 @Author: as
 @Date: Creat in 17:51 2022/2/13
 @Description: user.go
*/


type one struct {
	i int
}

func (o *one)try(){
	fmt.Println(o.i)
}


func TestAddTask(t *testing.T) {
	SetPro()
	wg:= sync.WaitGroup{}

	for i:=0;i<10;i++{
		wg.Add(1)
		o:=one{i: i}
		go AddTask(func() {
			o.try()
			wg.Done()
		})

	}
	wg.Wait()
}