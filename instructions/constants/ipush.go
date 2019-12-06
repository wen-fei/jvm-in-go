package constants

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

type BIPUSH struct {
	// push byte
	val int8
}

type SIPUSH struct {
	// push int
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
