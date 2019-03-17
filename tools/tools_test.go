package tools

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomStringX(t *testing.T) {
	rs := GenerateRandomStringX(5, 8, true, true, true, false, false, false)

	t.Logf("随机字符串：%v", rs)
}

func TestStrToInt(t *testing.T) {

	n1, errT := StringToInt("12")

	if errT == nil && n1 != 12 {
		t.Errorf("测试失败：n1为%v（预期值：%v）", n1, 12)
	}

	n2, errT := StringToInt("012")

	if errT == nil && n2 != 12 {
		t.Fatalf("测试失败：n2为%v（预期值：%v）", n2, 12)
	}

	n3, errT := StringToInt("ABZ")

	if errT == nil {
		t.Errorf("测试失败：errT为nil（预期应不为nil），n3为%v", n3)
	}

}

func TestStrToIntParallel(t *testing.T) {

	for i := 0; i < 1000000; i++ {
		n1 := rand.Intn(500)

		n2, errT := StringToInt(IntToString(n1))

		if errT != nil {
			t.Fatalf("测试失败：n1的值为%v，n2的值为%v, errT为%v", n1, n2, errT)
		}

		if n1 != n2 {
			t.Fatalf("测试失败：n1的值为%v，n2的值为%v", n1, n2)
		}
	}
}

func Test001(t *testing.T) {
	// t.Parallel()

	for i := 0; i < 5; i++ {
		t.Run("并发测试", TestStrToIntParallel)
	}
}

func Test002(t *testing.T) {
	// t.Parallel()

	for i := 0; i < 5; i++ {
		t.Run("并发测试", TestStrToIntParallel)
	}
}
