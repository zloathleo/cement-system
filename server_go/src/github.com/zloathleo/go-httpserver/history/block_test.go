package history

import (
	"testing"
)

// 单元测试
// 测试全局函数，以TestFunction命名
// 测试类成员函数，以TestClass_Function命名
func TestBlockToBytes(t *testing.T) {
	var dass = make([]*Das ,0)

	for i:=0 ;i<600;i++ {
		das := &Das{
			float32(i),
			1538297216,
		}
		dass = append(dass, das)
	}
	datas := blockToBytes(1538296642,99,dass)

	blockFromBytes(datas)
}
