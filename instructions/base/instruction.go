package base

import "github.com/jvm-in-go/rtda"

// 很多指令操作数类似，因此定义一个base接口，类似Java中的抽象类
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 存储和加载指令，根据索引存取局部变量表
type Index8Intstruction struct {
	Index int
}

func (self *Index8Intstruction) FetchOperands(reader *BytecodeReader) {
	self.Index = int(reader.ReadInt8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt16())
}
