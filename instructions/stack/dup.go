package stack

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

type DUP struct {
	base.NoOperandsInstruction
}

type DUP_X1 struct {
	base.NoOperandsInstruction
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

type DUP2 struct {
	base.NoOperandsInstruction
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

// cop single val on the top of stack
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
