/*
 add by jinchaozhang. 2021/4/22
*/
package alg_test_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

//[0,1)的平方
func xPower2() float64 {
	return math.Max(rand.Float64(), rand.Float64())
}

func TestXPower2(t *testing.T) {
	maxTimes := 1000000
	count := 0
	num := 0.3
	for i := 0; i < maxTimes; i++ {
		if xPower2() < num {
			count++
		}
	}
	fmt.Println(float64(count) / float64(maxTimes))
	fmt.Println(math.Pow(num, 2))
}

type randBox struct {
	Min int
	Max int
}

//[min,max]等概率返回
func (box *randBox) randPREqual() int {
	return box.Min + rand.Intn(box.Max-box.Min+1)
}

//根据上述[min,max]等概率，进行0，1 等概率返回
func (box *randBox) randPREqualBit() int {
	size := box.Max - box.Min + 1
	mid := size / 2
	odd := size&1 != 0
	var ans int
	for {
		ans = box.randPREqual() - box.Min
		if odd && ans == mid {
			ans = 0
		} else {
			break
		}
	}
	if ans < mid {
		return 0
	}
	return 1
}

//[from,to] 等概率返回
func (box *randBox) randNumber(from, to int) int {
	if from == to {
		return from
	}
	//[0,rangeNum] 2机制算概率 取全为1 大于等于range number 二进制位数
	rangeNum := to - from
	bitNums := 1
	//rangenum 大于最大的二进制位数需要进位
	for num := 1; rangeNum > num; num = num<<1 + 1 {
		bitNums++
	}
	num := 0
	for {
		for maxoffsetBit := bitNums - 1; maxoffsetBit >= 0; maxoffsetBit-- {
			num += box.randPREqualBit() << maxoffsetBit
		}
		//大于rangenum 需要重新分配（等概率）
		if num > rangeNum {
			num = 0
		} else {
			break
		}
	}

	return from + num
}

func TestBoxRandom(t *testing.T) {

	fmt.Println("start")
	box := &randBox{12, 16}
	randFrom := 2
	randEnd := 13
	arrLength := randEnd - randFrom + 1
	arr := make([]int, arrLength)
	maxTimes := 10000000
	for i := 0; i < maxTimes; i++ {
		num := box.randNumber(randFrom, randEnd)

		arr[num-2] = arr[num-2] + 1
	}
	for i, val := range arr {
		fs := fmt.Sprintf("num:%d,count:%d,pb:%f", i+2, val, float64(val)/float64(maxTimes))
		fmt.Println(fs)
	}
	fmt.Println("done")
}

func npr() int {
	if rand.Float64() < 0.8 {
		return 0
	}
	return 1
}

//不等概率函数 转 等概率 随机只取 0，1 和1，0 两种组合情况 两者概率是一样的 等概率
func pr() int {
	first := 0
	for {
		first = npr()

		if first != npr() {
			break
		}
	}
	return first
}

func TestNonePrToPrRandom(t *testing.T) {
	maxTimes := 1000000
	count := 0
	for i := 0; i < maxTimes; i++ {
		if pr() == 1 {
			count++
		}
	}
	fmt.Println(float64(count) / float64(maxTimes))
}
