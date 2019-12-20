package heap

type SymRef struct {
	// 运行时常量池指针
	cp        *ConstantPool
	className string
	class     *Class
}

// 类符号引用解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		// 如果类符号引用已经解析，直接返回类指针，否则调用resolveClassRef
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
