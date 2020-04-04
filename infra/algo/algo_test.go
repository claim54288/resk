package algo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleRand(t *testing.T) {
	FotTest("简单随机算法", t, SimpleRand)
}

func TestBeforeShuffle(t *testing.T) {
	FotTest("先洗牌算法", t, BeforeShuffle)
}

func TestDoubleRandom(t *testing.T) {
	FotTest("二次随机算法", t, DoubleRandom)
}

func TestDoubleAverage(t *testing.T) {
	FotTest("二倍随机算法", t, DoubleAverage)
}

func FotTest(messge string, t *testing.T, fn func(count, amount int64) int64) {
	count, amount := int64(10), int64(100*100)
	remain := amount
	sum := int64(0) //用来验证总金额的
	for i := int64(0); i < count; i++ {
		x := fn(count-i, remain)
		remain -= x
		sum += x
	}
	Convey(messge, t, func() {
		So(sum, ShouldEqual, amount)
	})
}
