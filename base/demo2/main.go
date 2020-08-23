package main

import (
	"fmt"
	"sync"
)

/**
生产者消费者模型


 */


// 生产者
func Producer(factor int ,out chan<- int,wg *sync.WaitGroup)  {
	defer func() {
		fmt.Println("生产者停止")
		wg.Done()
	}()
	
	i:=0
	for true {
		tmp:=i*factor
		fmt.Println("生产数据",tmp)
		out<-tmp
		i++
	}
}

// 消费者
func Consumer(in <-chan int,wg *sync.WaitGroup)  {
	defer func() {
		fmt.Println("消费者停止")
		wg.Done()
	}()
	
	for i:=range in{
		fmt.Println("消费数据:",i)
	}
}


func main() {
	var  wg sync.WaitGroup

	var q=make(chan int,10)
	
	// 启动后生产者
	wg.Add(1)
	go Producer(5,q,&wg)
	
	// 启动消费者
	wg.Add(3)
	go Consumer(q,&wg)
	go Consumer(q,&wg)
	go Consumer(q,&wg)
	
	// 等待完事
	wg.Wait()
	
}
