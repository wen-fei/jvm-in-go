package comparisons

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func _dcmp(self *rtda.Frame, flag bool) {
	stack := self.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}
