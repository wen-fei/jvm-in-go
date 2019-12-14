package rtda

import (
	"github.com/jvm-in-go/rtda/heap"
	"math"
)

// local vars table
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// float可以先转成int类型，再按int变量来处理
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self[index].num)
	return math.Float64frombits(bits)
}

// long变量需要拆成两个int变量
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)>>32 | int64(low)
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}
