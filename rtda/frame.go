package rtda

import "github.com/jvm-in-go/rtda/heap"

type Frame struct {
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 操作数栈指针
	operandStack *OperandStack
	method       *heap.Method
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}
