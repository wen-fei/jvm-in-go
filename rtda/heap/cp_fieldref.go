package heap

import "github.com/jvm-in-go/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 字段符号引用解析
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolvedFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolvedFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {
	// 查找当前类
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	// 递归查找接口
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	// 查找超类
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
