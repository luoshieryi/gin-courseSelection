package service

import (
	"sync"
	"testing"
)

/*
 @Author: as
 @Date: Creat in 16:48 2022/2/12
 @Description: user.go
*/
var c=courseInfo{
	cap:    10,
	left:   10,
	leftCh: make(chan int),
	done:   make(chan struct{}),
	close:  make(chan struct{}),
	lock:   &sync.Mutex{},
}

func BenchmarkBookCourse(t *testing.B) {

	c.monitor()
	for i:=0;i<t.N;i++{
		//fmt.Println(c.getLeft())
		c.bookOne("test","1")
	}
}