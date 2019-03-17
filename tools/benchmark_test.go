package tools

import (
	"testing"
)

func BenchmarkCalPi(b *testing.B) {
	for i := 1; i < b.N; i++ {
		rs := CalPi(i)
		b.Logf("Pi值：%v", rs)
	}
}

func BenchmarkCalPiX(b *testing.B) {
	for i := 1; i < b.N; i++ {
		rs := CalPiX(i)
		b.Logf("Pi值：%v", rs)
	}
}

func BenchmarkFibo38(b *testing.B) {
	rs := Fibonacci(38)
	b.Logf("斐波那契38结果值：%v", rs)
}

func BenchmarkFibo48(b *testing.B) {
	rs := Fibonacci(48)
	b.Logf("斐波那契48结果值：%v", rs)
}
