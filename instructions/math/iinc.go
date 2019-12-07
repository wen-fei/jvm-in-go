package math

import "github.com/jvm-in-go/rtda"

// iinc指令给局部变量表中的int变量增加常量值
type IINC struct {
	Index uint
	Const int32
}

// read int from local stack, add index, then push into stack again
func (self *IINC) FetchOperands(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
