package main

import (
	"math/rand"
	"time"
	t "tools"
)

// Request1 是自定义的结构类型
// 用于发送请求时传递所需的所有数据
type Request1 struct {
	ID       int           // 本次请求的编码，是随机产生的整数
	Count    int           // 计算多少次
	Response chan []string // 用于存放请求结果的通道
}

// 令牌池
var poolG chan int

// 缓存请求的通道
var requestChannelG chan Request1

// 发送请求并等待请求执行结果的函数
func sendRequest(countA int) {
	idT := rand.Intn(10000)

	responseChanT := make(chan []string, 1)

	defer close(responseChanT)

	requestChannelG <- Request1{ID: idT, Count: countA, Response: responseChanT}

	responseT := <-responseChanT

	t.Printfln("goroutine（ID: %v, Count: %v）结果: %#v", idT, countA, responseT)

}

// 处理具体每个请求的函数
func doRequest(requestA Request1) {
	resultT := 0
	for i := 0; i < requestA.Count; i++ {
		resultT = i + requestA.Count
	}

	requestA.Response <- []string{"请求成功", t.IntToString(resultT)}

	poolG <- 0
}

// 处理请求队列（通道）的函数
// 根据令牌池是否有剩余令牌来决定是否启动处理具体请求的goroutine
func processRequests() {
	for {
		requestT := <-requestChannelG

		if len(poolG) > 0 {
			<-poolG
			go doRequest(requestT)
		} else {
			requestT.Response <- []string{"请求失败", "goroutine池已满"}
		}

	}
}

func main() {

	// 为令牌池分配容量，容量虽然为10
	// 但可用令牌是根据后面放入的令牌个数来确定的
	poolG = make(chan int, 10)

	// 为请求缓冲通道分配100个的容量
	// 如果缓冲通道内超过100个请求，下一个请求将被阻塞
	requestChannelG = make(chan Request1, 100)

	// 确保通道退出时会被关闭
	defer close(requestChannelG)
	defer close(poolG)

	// 启动请求处理主goroutine
	go processRequests()

	// 在令牌池中放入5块令牌
	for i := 0; i < 5; i++ {
		poolG <- 0
	}

	// 模拟发送20个请求
	for i := 0; i < 20; i++ {
		go sendRequest(1000)
	}

	// 主线程死循环以保证处理请求的goroutine一直执行
	// 需要用Ctrl+C键来退出程序运行
	for {
		time.Sleep(100 * time.Millisecond)
	}

}
