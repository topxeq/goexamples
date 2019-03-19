package main

import (
	"runtime"
	"time"
	t "tools"
)

var goroutineCount int // 用于设置goroutine数量

var resultBuffer1 chan float64 // 用于放置各个addRoutine1计算的结果
var resultBuffer2 chan int     // 用于放置各个addRoutine2计算的结果

// 多goroutine计算浮点数累加和的单个goroutine
// countA代表本goroutine需要执行多少次计算
func addRoutine1(countA int) {
	sumT := 0.0

	for i := 0; i < countA; i++ {
		sumT += 1.1
	}

	// 将结果写入对应的通道中
	resultBuffer1 <- sumT
}

// 多goroutine计算整数累加和的单个goroutine
// countA代表本goroutine需要执行多少次计算
func addRoutine2(countA int) {
	sumT := 0

	for i := 0; i < countA; i++ {
		sumT += 7
	}

	// 将结果写入对应的通道中
	resultBuffer2 <- sumT
}

// 调用多个addRoutine1函数和多个addRoutine2实现各自累加和并将两个累加和求相除结果的函数
func addByGoroutine(countA int) float64 {
	sumT1 := 0.0
	sumT2 := 0

	// lenT是每个goroutine需要计算的次数
	lenT := countA / goroutineCount

	// leftT是平均非给每个goroutine后还剩余需要计算的次数
	leftT := countA - (countA/goroutineCount)*goroutineCount

	// 第一个goroutine将多计算leftT次，即lenT+leftT次
	// addRoutine1和addRoutine2都将被运行成同样个数的goroutine
	// 各自生成goroutineCount个goroutine
	go addRoutine1(lenT + leftT)
	go addRoutine2(lenT + leftT)

	// 其他goroutine将计算lenT次
	for i := 1; i < goroutineCount; i++ {
		go addRoutine1(lenT)
		go addRoutine2(lenT)
	}

	// 从通道循环读取resultBuffer1或resultBuffer2中的值
	// 直到读满足够的个数（应为2 * goroutineCount个）

	var tmpF float64
	var tmpC int

	timeoutFlag := false

	for i := 0; i < goroutineCount*2; i++ {
		select {
		case tmpF = <-resultBuffer1:
			sumT1 += tmpF
		case tmpC = <-resultBuffer2:
			sumT2 += tmpC
		case <-time.After(3 * time.Second):
			timeoutFlag = true
		}

		if timeoutFlag {
			return 0.0
		}
	}

	// 返回最终的计算结果，为了计算类型一致，需要强制类型转换
	return float64(sumT2) / sumT1
}

func main() {

	// 计算的次数
	times := 50000000000

	// 获取实际CPU核数
	cpuCores := runtime.NumCPU()
	t.Printfln("CPU核数: %v", cpuCores)

	// goroutine个数设为可用CPU核数
	goroutineCount = cpuCores

	// 结果缓冲区大小与goroutine个数应相等，以便接受足够个数的计算结果
	resultBuffer1 = make(chan float64, goroutineCount)
	resultBuffer2 = make(chan int, goroutineCount)

	startTime := time.Now()

	result := addByGoroutine(times)

	endTime := time.Now()

	// 别忘了关闭两个通道
	close(resultBuffer1)
	close(resultBuffer2)

	t.Printfln("计算结果: %v", result)

	t.Printfln("计算时长: %v", endTime.Sub(startTime))

}
