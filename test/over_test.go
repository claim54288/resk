package test

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

//基准测试代码，Benchmark开头，第一个字母B大写
//testing.B作为参数
func BenchmarkUpdateForLock(b *testing.B) {
	g := GoodsSigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.ReaminQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(1000000)
	_, err := db.Insert(g)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		UpdateForLock(g.Goods)
	}
}

//无符号字段类型+直接更新
func BenchmarkUpdateForUnsigned(b *testing.B) {
	g := GoodsUnsigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.ReaminQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(1000000)
	_, err := db.Insert(g)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		UpdateForUnsigned(g.Goods)
	}
}

//乐观锁的基准测试
func BenchmarkUpdateDorOptimistic(b *testing.B) {
	g := GoodsSigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.ReaminQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(1000000)
	_, err := db.Insert(g)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		UpdateDorOptimistic(g.Goods)
	}
}

//乐观锁+无符号的基准测试
func BenchmarkUpdateDorOptimisticAndUnsigned(b *testing.B) {
	g := GoodsUnsigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.ReaminQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(1000000)
	_, err := db.Insert(g)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		UpdateDorOptimisticAndUnsigned(g.Goods)
	}
}
