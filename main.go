package main

import (
	"errors"
	"fmt"
)

// 利用byte切片充当位图
type Bitmap []byte

// 一个byte 8个bit，切片长度除以8后+1
func NewBitmap(max uint) *Bitmap {
	length := max>>3 + 1
	sli := make(Bitmap, length)
	return &sli
}

// 添加元素 利用『或』使其位为1
func (bm *Bitmap) AddItem(item uint) error {
	if item > bm.Max() {
		return errors.New("item too large")
	}
	index := item >> 3 // item / 8
	pos := item & 0x7  // item % 8
	// fmt.Printf("index=%d, pos=%d\n", index, pos)
	// 这里是低位开始计数，比如十进制的3 在位图里为 0000 1000
	// 包含0这个数，位图里表示 0000 0001
	(*bm)[index] |= 1 << uint(pos)
	return nil
}

// 判断是否存在，利用『与』判断该位是否为1
func (bm *Bitmap) Exists(item uint) bool {
	index := item >> 3
	pos := item & 0x7
	return (*bm)[index]&(1<<uint(pos)) != 0
}

// 删除项 利用『取反 再 与』使其位为0
func (bm *Bitmap) DelItem(item uint) error {
	if item > bm.Max() {
		return errors.New("item too large")
	}
	index := item >> 3
	pos := item & 0x7
	(*bm)[index] = (*bm)[index] & ^(1 << uint(pos))
	return nil
}

// 当前位图最大支持多大数
func (bm *Bitmap) Max() uint {
	return uint(len(*bm)*8) - 1
}

func (bm *Bitmap) String() string {
	return fmt.Sprint(*bm)
}

// 是否全是0
func (bm *Bitmap) IsAllZero() bool {
	leng := len(*bm)
	num := int64(0)
	for _, v := range *bm {
		num++
		if v != 0 {
			return false
		}
	}
	if int64(leng) == num {
		return true
	}
	return false
}

// 是否全是1
func (bm *Bitmap) IsAllOne() bool {
	leng := len(*bm)
	num := int64(0)
	for _, v := range *bm {
		num++
		if v != 255 {
			return false
		}
	}
	if int64(leng) == num {
		return true
	}
	return false
}

// 获取位置信息
func (bm *Bitmap) GetData() []int32 {
	ret := make([]int32, 0)
	for i := int32(1); i <= int32(bm.Max()); i++ {
		if bm.Exists(uint(i)) {
			ret = append(ret, i)
		}
	}
	return ret
}

// 测试
func main() {
	bm := NewBitmap(15)
	fmt.Println(bm.IsAllZero())
	bm.AddItem(0)
	fmt.Println(bm)

	for i := uint(0); i <= 15; i++ {
		bm.AddItem(i)
	}
	fmt.Println(bm.Exists(7), bm.Exists(10))
	fmt.Println(bm)

	fmt.Println(bm.IsAllOne())

	fmt.Println(bm)

	bm.DelItem(7)
	fmt.Println(bm.Exists(7), bm.Exists(10))

	fmt.Println(bm.GetData())
}
