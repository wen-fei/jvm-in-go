package math

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

// int 左移
type ISHL struct {
	base.NoOperandsInstruction
}

// int算术右位移
type ISHR struct {
	base.NoOperandsInstruction
}

// int逻辑右位移
type IUSHR struct {
	base.NoOperandsInstruction
}

// long左位移
type LSHL struct {
	base.NoOperandsInstruction
}

// long算术右位移
type LSHR struct {
	base.NoOperandsInstruction
}

// long逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f //Go语言位移操作符右侧必须是无符号整数，同事int只有32位，取前5byte就够了
	result := v1 << s
	stack.PushInt(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f // long64位，所以取前6个比特
	result := v1 >> s
	stack.PushLong(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// Go没有 >>> 运算符，为了达到无符号位移的目的，需要先把v1转成无符号整数，位移操作后再转回有符号整数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
