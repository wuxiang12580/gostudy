package main

import (
	"fmt"
	real2 "retiever/real"
	"time"
)

type Retiever interface {
	GetFrom(url string) string
}


func download(r Retiever)string  {
	return r.GetFrom("http://www.imooc.com");
}

//定义接口Retiever，里面有一个get方法，在mock.go 和 real.go这两个文件里分别实现了get方法，并重写了get方法

func main() {
	var r Retiever
	//r = mock.Retiever{"hello world"}
	//fmt.Println(r)

	r = &real2.Retiever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	//fmt.Println(download(r))
	fmt.Printf("%T %v\n",r,r)
}
