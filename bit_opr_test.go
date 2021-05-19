/*
 add by jinchaozhang. 2021/4/26
*/
package alg_test_test

import (
	"fmt"
	"testing"
)

type BitMap struct {
	bitSets []int64
}

func newBitMap(max int) *BitMap {
	bitMap := new(BitMap)
	bitMap.bitSets = make([]int64, (max+64)>>6)
	return bitMap
}

func (b *BitMap) add(num int) {
	b.bitSets[num>>6] |= 1 << (num & 63)
}

func (b *BitMap) delete(num int) {
	b.bitSets[num>>6] &= ^(1 << (num & 63))
}
func (b *BitMap) contains(num int) bool {
	return b.bitSets[num>>6]&(1<<(num&63)) != 0
}

func TestBitMap(t *testing.T) {
	b := newBitMap(125)
	b.add(3)
	b.add(40)
	b.add(67)
	b.add(120)
	b.add(112)
	b.delete(120)
	fmt.Println(b.contains(112))
	fmt.Println(b.contains(115))
	fmt.Println(len(b.bitSets))
	fmt.Println(fmt.Sprintf("%b", b.bitSets[0]))
	fmt.Println(fmt.Sprintf("%b", b.bitSets[1]))
}

func add(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}
func negNum(n int) int {
	return add(^n, 1)
}
func minus(a, b int) int {
	return add(a, negNum(b))
}

func multi(a, b int) int {
	res := 0
	for b != 0 {
		if b&1 != 0 {
			res = add(res, a)
		}
		a = a << 1
		b = int(uint(b) >> 1)
	}

	return res
}

func div(a, b int32) int {
	res := 0
	var x, y int32
	x = a
	y = b
	if a < 0 {
		x = int32(negNum(int(a)))
	}
	if b < 0 {
		y = int32(negNum(int(b)))
	}
	for i := 30; i >= 0; i = minus(i, 1) {
		if (x >> i) >= y {
			res |= (1 << i)
			x = int32(minus(int(x), int(y<<i)))
		}
	}

	if (a < 0) == (b < 0) {
		return res
	}
	return negNum(res)
}

func TestBitOpr(t *testing.T) {
	sum := add(5, 3)
	fmt.Println(sum)
	sub := minus(5, 3)
	fmt.Println(sub)
	m := multi(5, -3)
	fmt.Println(m)
	d := div(9, -3)
	fmt.Println(d)
}
