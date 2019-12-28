package control

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

type RETURN struct {
	// return void from method
	base.NoOperandsInstruction
}

type ARETURN struct {
	// return reference from method
	base.NoOperandsInstruction
}

type DRETURN struct {
	// return double from method
	base.NoOperandsInstruction
}

type FRETURN struct {
	// return float from method
	base.NoOperandsInstruction
}

type IRETURN struct {
	// return int from method
	base.NoOperandsInstruction
}

type LRETURN struct {
	// return long from method
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(retVal)
}
