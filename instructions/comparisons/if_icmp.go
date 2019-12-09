package comparisons

import (
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/rtda"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
