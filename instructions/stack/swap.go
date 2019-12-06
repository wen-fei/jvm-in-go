package stack

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

// swap the top two operand stack values
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
