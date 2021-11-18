package bisect_test

import (
	"bisect"
	"testing"
)

func TestStop_block(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := 4

	ans := bisect.BisectLeft(a, key)
	// 期待する出力とは違うansが得られたら t.Errorfメソッドでテストが失敗するようにする
	if ans != 3 {
		t.Errorf("get wrong ans : %d", ans)
	}
}
