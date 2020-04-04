package algo

import (
	"math/rand"
)

const min = int64(1)

//简单随机算法
//参数：红包的数量，红包金额
//使用float32 float64计算的时候会丢失精度，所以定义时候乘个100，用int64
//金额单位为分，1块= 100分
func SimpleRand(count, amount int64) int64 {
	//当红包数量剩余一个的时候，就直接返回剩余金额
	if count == 1 {
		return amount
	}
	//计算最大可调度金额
	max := amount - min*count
	x := rand.Int63n(max) + min
	return x
}
