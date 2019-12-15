package heap

type SymRef struct {
	// 运行时常量池指针
	cp        *ConstantPool
	className string
	class     *Class
}
