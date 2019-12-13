package extended

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

// branch if reference is null
type IFNULL struct {
	base.BranchInstruction
}

// branch if reference not null
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
