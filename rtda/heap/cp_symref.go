package heap

import "github.com/jvm-in-go/classfile"

type SymRef struct {
	// 运行时常量池指针
	cp        *classfile.ConstantPool
	className string
	class     *Class
}
